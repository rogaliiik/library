package app

import (
	"log"

	config "github.com/rogaliiik/library/config"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	slog, err := NewLogger(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Logger was inited")
}
