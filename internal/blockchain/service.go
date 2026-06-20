package blockchain

import (
	"blockchain/internal/config"
	"blockchain/services/pkg/certificate"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/skip2/go-qrcode"
)

type BlockChainService interface {
	// block chain methods
	IssueCertificate(ctx context.Context, req IssueCertificateRequest) (*IssueCertificateResponse, error)
	VerifyCertificate(ctx context.Context, pdfHash [32]byte) (*VerifyCertificateResponse, error)
	RevokeCertificate(ctx context.Context, pdfHash [32]byte) error

	// other methods
	GenerateQRCode(content string) ([]byte, error)
}

type blockChainService struct {
	cnf      *config.Config
	client   *ethclient.Client
	contract *certificate.Certificate
	repo     ProjectRepository
}

func NewBlockChainService(cnf *config.Config, repo ProjectRepository) BlockChainService {
	client, err := ethclient.Dial(cnf.SepoliaRPCURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(cnf.ContractAddress)

	contractInstance, err := certificate.NewCertificate(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &blockChainService{
		cnf:      cnf,
		client:   client,
		contract: contractInstance,
		repo:     repo,
	}
}

func (bcc *blockChainService) IssueCertificate(ctx context.Context, req IssueCertificateRequest) (*IssueCertificateResponse, error) {
	auth, err := bcc.getAuth(ctx)
	if err != nil {
		return nil, err
	}

	// issue on blockchain
	tx, err := bcc.contract.IssueCertificate(auth, req.PdfHash, req.RecipientName, req.CourseName, req.Grade, req.IssuingAuthority)
	if err != nil {
		return nil, err
	}

	receipt, err := bind.WaitMined(ctx, bcc.client, tx)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, errors.New("issue certificate transaction failed")
	}

	// store in db
	err = bcc.repo.CreateCertificate(ctx, CreateCertificateParams{
		InstituteID:      req.InstituteID,
		CertificateHash:  common.Bytes2Hex(req.PdfHash[:]),
		RecipientName:    req.RecipientName,
		CourseName:       req.CourseName,
		Grade:            req.Grade,
		IssuingAuthority: req.IssuingAuthority,
		BlockchainTxHash: tx.Hash().Hex(),
		IssuedAt:         time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	log.Println("tx hash: ", tx.Hash().Hex())

	// generate qr code
	verifyURL := fmt.Sprintf("http://localhost:8080/certificates/verify/0x%s", common.Bytes2Hex(req.PdfHash[:]))
	qrCode, err := bcc.GenerateQRCode(verifyURL)
	if err != nil {
		return nil, err
	}

	return &IssueCertificateResponse{
		QRCode:           qrCode,
		VerifyURL:        verifyURL,
		CertificateHash:  common.Bytes2Hex(req.PdfHash[:]),
		BlockchainTxHash: tx.Hash().Hex(),
		RecipientName:    req.RecipientName,
		CourseName:       req.CourseName,
		Grade:            req.Grade,
		IssuedAt:         time.Now().UTC(),
	}, nil
}

func (bcc *blockChainService) VerifyCertificate(ctx context.Context, pdfHash [32]byte) (*VerifyCertificateResponse, error) {
	result, err := bcc.contract.VerifyCertificate(&bind.CallOpts{Context: ctx}, pdfHash)
	if err != nil {
		isValid, validErr := bcc.contract.IsCertificateValid(&bind.CallOpts{Context: ctx}, pdfHash)
		if validErr == nil && !isValid {
			return &VerifyCertificateResponse{
				IsValid: false,
			}, nil
		}
		return nil, err
	}

	return &VerifyCertificateResponse{
		RecipientName:    result.RecipientName,
		CourseName:       result.CourseName,
		Grade:            result.Grade,
		IssuingAuthority: result.IssuingAuthority,
		IssueDate:        result.IssueDate.Uint64(),
		IsValid:          result.IsValid,
	}, nil
}

func (bcc *blockChainService) RevokeCertificate(ctx context.Context, pdfHash [32]byte) error {
	auth, err := bcc.getAuth(ctx)
	if err != nil {
		return err
	}

	tx, err := bcc.contract.RevokeCertificate(auth, pdfHash)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, bcc.client, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return errors.New("revoke certificate transaction failed")
	}

	err = bcc.repo.RevokeCertificate(ctx, common.Bytes2Hex(pdfHash[:]))
	if err != nil {
		return err
	}

	log.Println("tx hash:", tx.Hash().Hex())

	return nil
}

func (bcc *blockChainService) getAuth(ctx context.Context) (*bind.TransactOpts, error) {
	// convert hex private key string into ECDSA private key
	privateKey, err := crypto.HexToECDSA(bcc.cnf.PrivateKey)
	if err != nil {
		return nil, err
	}

	// extract public key from private key
	publicKey := privateKey.Public()

	// cast generic public key into ECDSA public key
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not ECDSA")
	}

	// derive ethereum wallet address from public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get next transaction nonce for wallet
	nonce, err := bcc.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	// fetch suggested gas price from network
	gasPrice, err := bcc.client.SuggestGasPrice(ctx)

	if err != nil {
		return nil, err
	}

	// create chain id object for sepolia
	chainID := big.NewInt(int64(bcc.cnf.ChainId))

	// create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		return nil, err
	}

	// set transaction nonce
	auth.Nonce = big.NewInt(int64(nonce))

	// amount of ETH to send with tx (0 for contract calls)
	auth.Value = big.NewInt(0)

	// maximum gas allowed for tx execution
	auth.GasLimit = uint64(300000)

	// gas price to pay validators
	auth.GasPrice = gasPrice

	return auth, nil
}

func (bcc *blockChainService) GenerateQRCode(content string) ([]byte, error) {
	return qrcode.Encode(
		content,
		qrcode.Medium,
		256,
	)
}
