package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		address: ":8080",
		db:      dbConfig{},
	}

	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("Server has failed to start: %v", err)
		os.Exit(1)
	}
}
