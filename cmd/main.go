package main

import (
	"log"

	"github.com/IbnuFarhanS/go-pinjaman-online/config"
	"github.com/IbnuFarhanS/go-pinjaman-online/server"
)

func main() {
	cfg := config.NewConfig()
	cfg.Load()

	server := server.NewServer()
	err := server.Init(cfg.ConnectionString)
	if err != nil {
		log.Fatalf("failed to initialize the server %v", err)
	}

	err = server.Start(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("failed")
	}
}
