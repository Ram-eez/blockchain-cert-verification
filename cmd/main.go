package main

import (
	"blockchain/bootstrap"
	"blockchain/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// get env vars and load config
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}
	cfg := config.LoadConfig()

	appServer, err := bootstrap.InitializeApplication(cfg)
	if err != nil {
		log.Fatal(err)
	}

	appServer.Start()

}
