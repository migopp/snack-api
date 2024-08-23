package main

import (
	"log"

	"github.com/migopp/snack-api/internal/api"
)

func main() {
	server := api.Server{
		IP:    "localhost",
		Port:  8000,
		Store: api.CreateMockStore(),
	}
	if err := server.Run(); err != nil {
		log.Fatal("Error:", err)
	}
}
