package blockchain

import (
	"blockchain/internal/config"
	"blockchain/services/pkg/certificate"
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChainService interface {
	IssueCertificate(ctx context.Context, certificateId [32]byte) error
	VerifyCertificate(ctx context.Context, certificateId [32]byte) error
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

func (bcc *blockChainService) IssueCertificate(ctx context.Context, certificateId [32]byte) error

func (bcc *blockChainService) VerifyCertificate(ctx context.Context, certificateId [32]byte) error

func (bcc *blockChainService) RevokeCertificate(ctx context.Context, certificateId [32]byte) error
