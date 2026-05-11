package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	// blockchain
	SepoliaRPCURL   string
	PrivateKey      string
	ContractAddress string
	ChainId         int

	// postgres
	PgUsername     string
	PgPassword     string
	PgHost         string
	PgPort         string
	PgDatabaseName string
	PgSSLMode      string
}

func LoadConfig() *Config {

	chainID, err := strconv.Atoi(os.Getenv("CHAIN_ID"))
	if err != nil {
		log.Fatal("invalid CHAIN_ID")
	}

	cfg := &Config{
		// blockchain
		SepoliaRPCURL:   mustEnv("SEPOLIA_RPC_URL"),
		PrivateKey:      mustEnv("PRIVATE_KEY"),
		ContractAddress: mustEnv("CONTRACT_ADDRESS"),
		ChainId:         chainID,

		// postgres
		PgUsername:     mustEnv("PG_USERNAME"),
		PgPassword:     mustEnv("PG_PASSWORD"),
		PgHost:         mustEnv("PG_HOST"),
		PgPort:         mustEnv("PG_PORT"),
		PgDatabaseName: mustEnv("PG_DATABASE_NAME"),
		PgSSLMode:      mustEnv("PG_SSL_MODE"),
	}

	return cfg
}

func mustEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("missing environment variable: %s", key)
	}

	return value
}
