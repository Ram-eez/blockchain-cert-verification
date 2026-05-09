package backend

import (
	"blockchain/services/pkg/certificate"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockChainClient struct {
	Client   *ethclient.Client
	Contract *certificate.Certificate
}

// connect to sepolia

func Init() *BlockChainClient {
	client, err := ethclient.Dial("RPC LINK")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("contract address here")

	contractInstance, err := certificate.NewCertificate(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &BlockChainClient{
		Client:   client,
		Contract: contractInstance,
	}
}
