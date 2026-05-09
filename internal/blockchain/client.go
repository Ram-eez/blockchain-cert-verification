package blockchain

import (
	"blockchain/services/pkg/certificate"
	"fmt"
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
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cc504c934b524590bdc598a83f0cdcfd")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x9501dD95F4f9f502570e579cc54D42af02E63d99")

	contractInstance, err := certificate.NewCertificate(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &BlockChainClient{
		Client:   client,
		Contract: contractInstance,
	}
}

func (bc *BlockChainClient) VerifyCertificate(certificateID [32]byte) error {

	result, err := bc.Contract.VerifyCertificate(
		nil,
		certificateID,
	)

	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
