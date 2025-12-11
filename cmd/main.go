package main

import (
	"log/slog"
	"os"

	"github.com/Nios-V/ecommerce/api/internal/config"
)

func main() {
	config.Load()

	api := application{
		config: config.AppConfig,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server has failed to start", "error", err)
		os.Exit(1)
	}
}
