package main

import (
	"log"

	"github.com/migopp/snack-api/internal/api"
)

func main() {
	// init db
	postgresStore, err := api.CreatePostgresStore()
	defer postgresStore.Close()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	err = postgresStore.CreateSnackerTable()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	// create server
	server := api.Server{
		IP:    "localhost",
		Port:  8000,
		Store: postgresStore,
	}

	// spin
	if err = server.Run(); err != nil {
		log.Fatal("Error: ", err)
	}
}
