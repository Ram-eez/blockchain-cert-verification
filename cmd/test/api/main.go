package main

import (
	"blockchain/internal/blockchain"
	"log"
)

func main() {

	bc := blockchain.Init()

	var certID [32]byte
	copy(certID[:], []byte("test-certificate"))

	err := bc.VerifyCertificate(certID)
	if err != nil {
		log.Fatal(err)
	}
}
