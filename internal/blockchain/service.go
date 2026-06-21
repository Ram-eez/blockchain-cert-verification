package blockchain

import (
	"blockchain/internal/config"
	"blockchain/internal/middleware"
	"blockchain/internal/utils"
	"blockchain/services/pkg/certificate"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrCertificateNotFound       = errors.New("certificate does not exist")
	ErrCertificateAlreadyRevoked = errors.New("certificate already revoked")
)

type BlockChainService interface {
	// block chain methods
	IssueCertificate(ctx context.Context, req IssueCertificateRequest) (*IssueCertificateResponse, error)
	VerifyCertificate(ctx context.Context, pdfHash [32]byte) (*VerifyCertificateResponse, error)
	RevokeCertificate(ctx context.Context, pdfHash [32]byte) error

	// other methods
	GenerateQRCode(content string) ([]byte, error)
	Login(ctx context.Context, req LoginRequest) (string, error)
	CreateInstitution(ctx context.Context, req CreateInstitutionRequest) (*CreateInstitutionResponse, error)
}

type blockChainService struct {
	cnf      *config.Config
	client   *ethclient.Client
	contract *certificate.Certificate
	repo     ProjectRepository
	jwt      middleware.Middleware
	txMu     sync.Mutex
}

func NewBlockChainService(cnf *config.Config, repo ProjectRepository, jwt middleware.Middleware) BlockChainService {
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
		jwt:      jwt,
	}
}

func (bcc *blockChainService) IssueCertificate(ctx context.Context, req IssueCertificateRequest) (*IssueCertificateResponse, error) {
	// prevent concurrent transaction submissions
	bcc.txMu.Lock()
	defer bcc.txMu.Unlock()

	// check if certificate already exists
	result, err := bcc.VerifyCertificate(ctx, req.PdfHash)
	if err != nil {
		return nil, err
	}

	if result.Exists {
		if result.IsValid {
			return nil, errors.New("certificate already issued")
		}

		return nil, errors.New(
			"certificate was revoked; issue a corrected certificate with a different pdf",
		)
	}

	// create authenticated signer
	auth, err := bcc.getAuth(ctx)
	if err != nil {
		log.Println("failed to get auth:", err)
		return nil, err
	}

	log.Println("issuing certificate on blockchain")
	log.Println("pdf hash:", common.Bytes2Hex(req.PdfHash[:]))

	// issue certificate on blockchain
	tx, err := bcc.contract.IssueCertificate(
		auth,
		req.PdfHash,
		req.RecipientName,
		req.CourseName,
		req.Grade,
		req.IssuingAuthority,
	)
	if err != nil {
		log.Println("failed to send transaction:", err)
		return nil, err
	}

	log.Println("submitted tx:", tx.Hash().Hex())

	// wait for transaction to be mined
	log.Println("waiting for mining")

	mineCtx, cancel := context.WithTimeout(
		ctx,
		2*time.Minute,
	)
	defer cancel()

	receipt, err := bind.WaitMined(
		mineCtx,
		bcc.client,
		tx,
	)
	if err != nil {
		log.Println("failed waiting for receipt:", err)
		return nil, err
	}

	log.Println("wait mined returned")
	log.Println("receipt status:", receipt.Status)

	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Println("transaction reverted")

		return nil,
			errors.New(
				"issue certificate transaction failed",
			)
	}

	log.Println("storing certificate in database")

	issuedAt := time.Now()

	// store certificate metadata
	err = bcc.repo.CreateCertificate(
		ctx,
		CreateCertificateParams{
			InstituteID:      req.InstituteID,
			CertificateHash:  common.Bytes2Hex(req.PdfHash[:]),
			RecipientName:    req.RecipientName,
			CourseName:       req.CourseName,
			Grade:            req.Grade,
			IssuingAuthority: req.IssuingAuthority,
			BlockchainTxHash: tx.Hash().Hex(),
			IssuedAt:         issuedAt,
		},
	)
	if err != nil {
		log.Println("failed to store certificate:", err)
		return nil, err
	}

	log.Println("certificate stored successfully")

	// generate verification url
	verifyURL := fmt.Sprintf(
		"%s/verify/%s",
		strings.TrimRight(
			bcc.cnf.BaseURL,
			"/",
		),
		common.Bytes2Hex(
			req.PdfHash[:],
		),
	)

	// generate qr code
	qrCode, err := bcc.GenerateQRCode(
		verifyURL,
	)
	if err != nil {
		log.Println("failed generating qr:", err)
		return nil, err
	}

	log.Println("generated qr code")

	txURL := fmt.Sprintf(
		"https://sepolia.etherscan.io/tx/%s",
		tx.Hash().Hex(),
	)

	// generate downloadable pdf receipt
	pdfBytes, err := utils.GenerateCertificateCard(
		utils.CertificateCard{
			RecipientName:    req.RecipientName,
			CourseName:       req.CourseName,
			Grade:            req.Grade,
			IssuingAuthority: req.IssuingAuthority,

			IssuedAt: issuedAt.Format(
				time.RFC1123,
			),

			CertificateHash: common.Bytes2Hex(
				req.PdfHash[:],
			),

			TransactionHash: tx.Hash().Hex(),

			VerifyURL: verifyURL,

			TransactionURL: txURL,

			QRCode: qrCode,
		},
	)
	if err != nil {
		log.Println(
			"failed generating pdf:",
			err,
		)

		return nil, err
	}

	log.Println("generated pdf receipt")

	return &IssueCertificateResponse{
		QRCodeBase64: base64.StdEncoding.EncodeToString(
			qrCode,
		),

		PDFBase64: base64.StdEncoding.EncodeToString(
			pdfBytes,
		),

		VerifyURL: verifyURL,

		CertificateHash: common.Bytes2Hex(
			req.PdfHash[:],
		),

		BlockchainTxHash: tx.Hash().Hex(),

		TransactionURL: txURL,

		RecipientName: req.RecipientName,

		CourseName: req.CourseName,

		Grade: req.Grade,

		IssuedAt: issuedAt,

		IssuingAuthority: req.IssuingAuthority,
	}, nil
}
func (bcc *blockChainService) VerifyCertificate(ctx context.Context, pdfHash [32]byte) (*VerifyCertificateResponse, error) {
	result, err := bcc.contract.VerifyCertificate(&bind.CallOpts{Context: ctx}, pdfHash)

	if err != nil {
		isValid, validErr := bcc.contract.IsCertificateValid(&bind.CallOpts{Context: ctx}, pdfHash)

		if validErr == nil && !isValid {
			return &VerifyCertificateResponse{Exists: false, IsValid: false}, nil
		}

		return nil, err
	}

	cert, err := bcc.repo.GetCertificateByHash(ctx, common.Bytes2Hex(pdfHash[:]))
	if err != nil {
		return nil, err
	}

	verifyURL := fmt.Sprintf("%s/verify/%s", strings.TrimRight(bcc.cnf.BaseURL, "/"), cert.CertificateHash)

	txURL := fmt.Sprintf("https://sepolia.etherscan.io/tx/%s", cert.BlockchainTxHash)

	return &VerifyCertificateResponse{
		Exists:           true,
		IsValid:          result.IsValid,
		IsRevoked:        cert.IsRevoked,
		RecipientName:    cert.RecipientName,
		CourseName:       cert.CourseName,
		Grade:            cert.Grade,
		IssuingAuthority: cert.IssuingAuthority,
		IssueDate:        result.IssueDate.Uint64(),
		IssuedAt:         cert.IssuedAt,
		CertificateHash:  cert.CertificateHash,
		BlockchainTxHash: cert.BlockchainTxHash,
		VerifyURL:        verifyURL,
		TransactionURL:   txURL,
	}, nil
}

func (bcc *blockChainService) RevokeCertificate(ctx context.Context, pdfHash [32]byte) error {
	certificateHash := common.Bytes2Hex(pdfHash[:])

	cert, err := bcc.repo.GetCertificateByHash(ctx, certificateHash)
	if err != nil {
		return ErrCertificateNotFound
	}

	if cert.IsRevoked {
		return ErrCertificateAlreadyRevoked
	}

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

	err = bcc.repo.RevokeCertificate(ctx, certificateHash)
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

	// fetch wallet balance
	balance, err := bcc.client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, err
	}

	log.Println("wallet:", fromAddress.Hex())
	log.Println("balance:", balance)

	// compare mined and pending nonces to detect mempool drift
	confirmedNonce, err := bcc.client.NonceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, err
	}

	pendingNonce, err := bcc.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	log.Println("confirmed nonce:", confirmedNonce)
	log.Println("pending nonce:", pendingNonce)

	if pendingNonce > confirmedNonce {
		log.Printf(
			"wallet has %d pending transactions",
			pendingNonce-confirmedNonce,
		)
	}

	// create chain id object for sepolia
	chainID := big.NewInt(int64(bcc.cnf.ChainId))

	// create authenticated transaction signer
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	// amount of ETH to send with tx (0 for contract calls)
	auth.Value = big.NewInt(0)

	// maximum gas allowed for tx execution
	auth.GasLimit = uint64(300000)

	// allow go-ethereum to determine nonce and fees automatically

	return auth, nil
}

func (bcc *blockChainService) GenerateQRCode(content string) ([]byte, error) {
	return qrcode.Encode(
		content,
		qrcode.Medium,
		256,
	)
}

func (bcc *blockChainService) Login(ctx context.Context, req LoginRequest) (string, error) {
	institute, err := bcc.repo.GetInstituteByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(institute.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := bcc.jwt.GenerateJWT(institute.ID, institute.Name)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (bcc *blockChainService) CreateInstitution(ctx context.Context, req CreateInstitutionRequest) (*CreateInstitutionResponse, error) {
	name := strings.TrimSpace(req.Name)
	email := strings.TrimSpace(strings.ToLower(req.Email))

	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	password := utils.GeneratePassword(12)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = bcc.repo.CreateInstitute(ctx, CreateInstituteModel{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return nil, err
	}

	return &CreateInstitutionResponse{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
