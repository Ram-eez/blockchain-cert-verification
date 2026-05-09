package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SepoliaRPCURL   string
	PrivateKey      string
	ContractAddress string
	ChainId         int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	chainID, err := strconv.Atoi(os.Getenv("CHAIN_ID"))
	if err != nil {
		log.Fatal("invalid chain id")
	}

	cfg := &Config{
		SepoliaRPCURL:   os.Getenv("SEPOLIA_RPC_URL"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		ChainId:         chainID,
	}

	return cfg
}
