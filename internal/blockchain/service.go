package blockchain

import (
	"blockchain/internal/config"
	"blockchain/services/pkg/certificate"
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChainService interface {
	IssueCertificate(ctx context.Context, req IssueCertificateRequest) error
	VerifyCertificate(ctx context.Context, certificateId [32]byte) (*VerifyCertificateResponse, error)
	RevokeCertificate(ctx context.Context, certificateId [32]byte) error
}

type blockChainService struct {
	cnf      *config.Config
	client   *ethclient.Client
	contract *certificate.Certificate
}

func NewBlockChainService(cnf *config.Config) BlockChainService {
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
	}
}

func (bcc *blockChainService) IssueCertificate(ctx context.Context, req IssueCertificateRequest) error {
	auth, err := bcc.getAuth()
	if err != nil {
		return err
	}

	tx, err := bcc.contract.IssueCertificate(auth, req.PdfHash, req.RecipientName, req.CourseName, req.Grade, req.IssuingAuthority)
	if err != nil {
		return err
	}

	// store later in db as well
	log.Println("tx hash: ", tx.Hash().Hex())

	return nil
}

func (bcc *blockChainService) VerifyCertificate(ctx context.Context, certificateId [32]byte) (*VerifyCertificateResponse, error) {
	result, err := bcc.contract.VerifyCertificate(&bind.CallOpts{Context: ctx}, certificateId)
	if err != nil {
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

func (bcc *blockChainService) RevokeCertificate(ctx context.Context, certificateId [32]byte) error {
	auth, err := bcc.getAuth()
	if err != nil {
		return err
	}

	tx, err := bcc.contract.RevokeCertificate(
		auth,
		certificateId,
	)

	if err != nil {
		return err
	}

	log.Println("tx hash:", tx.Hash().Hex())

	return nil
}

func (bcc *blockChainService) getAuth() (*bind.TransactOpts, error) {
	// convert hex private key string into ECDSA private key
	privateKey, err := crypto.HexToECDSA(bcc.cnf.PrivateKey)
	if err != nil {
		return nil, err
	}

	// extract public key from private key
	publicKey := privateKey.Public()

	// cast generic public key into ECDSA public key
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)

	// derive ethereum wallet address from public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// get next transaction nonce for wallet
	nonce, err := bcc.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	// fetch suggested gas price from network
	gasPrice, err := bcc.client.SuggestGasPrice(context.Background())

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
