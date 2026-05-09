package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SepoliaRPCURL   string
	PrivateKey      string
	ContractAddress string
	ChainId         string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}

	cfg := &Config{
		SepoliaRPCURL:   os.Getenv("SEPOLIA_RPC_URL"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		ChainId:         os.Getenv("CHAIN_ID"),
	}

	return cfg
}
