package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/rogaliiik/library/config"
	v1 "github.com/rogaliiik/library/internal/handler/http/v1"
	"github.com/rogaliiik/library/internal/repository"
	"github.com/rogaliiik/library/internal/service"
	"github.com/rogaliiik/library/pkg/http_server"
)

// @title LibraryAPI
// @version 1.0
// @description API for Library app.

// @securityDefinitions.apikey UserAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := NewLogger(cfg.Log.Level)
	if err != nil {
		log.Fatal(err)
	}
	logger.Debug("Logger was inited")
	logger.Debug("Parsed config", slog.Any("config", cfg))

	db, err := connectPostgres(cfg.Postgres)
	if err != nil {
		logger.Error("Postgres connection error", slog.Any("message", err.Error()))
		os.Exit(1)
	}
	defer db.Close()
	logger.Debug("Postgres was inited")

	repos := repository.NewRepository(db)
	logger.Debug("Repositories was inited")

	services := service.NewServices(repos, cfg.Salt, cfg.SignKey, cfg.TokenTTL)
	logger.Debug("Services was inited")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := v1.NewHandler(ctx, services, logger).InitRoutes()
	logger.Debug("Handlers was inited")

	httpServer := http_server.New(router, cfg.Http.Port)
	logger.Debug(fmt.Sprintf("Http server was started at port: %s", cfg.Http.Port))

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)

	select {
	case <-notify:
		logger.Debug("Stop signal received")
	case <-httpServer.Notify():
		logger.Error("Server running error", slog.Any("message", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		logger.Error("Error while shutting down the server", slog.Any("message", err.Error()))
	}
}

func connectPostgres(cfg config.Postgres) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DbName, cfg.Username, cfg.Password)
	return sql.Open("postgres", dsn)
}
