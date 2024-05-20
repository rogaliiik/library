package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/rogaliiik/library/config"
	"github.com/rogaliiik/library/internal/repository"
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

	db, err := connectPostgres(cfg.Postgres)
	slog.Info("Postgres was inited")

	_ = repository.NewRepository(db)
	slog.Info("Repositories was inited")

}

func connectPostgres(cfg config.Postgres) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DbName, cfg.Username, cfg.Password)
	return sql.Open("postgres", dsn)
}
