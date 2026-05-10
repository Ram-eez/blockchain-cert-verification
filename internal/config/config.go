package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SepoliaRPCURL   string
	PrivateKey      string
	ContractAddress string
	ChainId         int
	PgUsername      string
	PgPassword      string
	PgHost          string
	PgPort          string
	PgDatabaseName  string
	PgSSLMode       string
}

func LoadConfig() *Config {
	envPath, err := findEnvFile()
	if err != nil {
		log.Fatal("failed to find .env")
	}
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("failed to load .env")
	}

	chainID, err := strconv.Atoi(os.Getenv("CHAIN_ID"))
	if err != nil {
		log.Fatal("invalid chain id")
	}

	cfg := &Config{
		// blockchain
		SepoliaRPCURL:   os.Getenv("SEPOLIA_RPC_URL"),
		PrivateKey:      os.Getenv("PRIVATE_KEY"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		ChainId:         chainID,
		// postgres
		PgUsername:     os.Getenv(""),
		PgPassword:     os.Getenv(""),
		PgHost:         os.Getenv(""),
		PgPort:         os.Getenv(""),
		PgDatabaseName: os.Getenv(""),
		PgSSLMode:      os.Getenv(""),
	}

	return cfg
}

func findEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		candidate := filepath.Join(currentDir, ".env")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			return "", os.ErrNotExist
		}
		currentDir = parent
	}
}
