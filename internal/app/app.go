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

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		logger.Fatal("Postgres connection error", err)
	}
	defer db.Close()
	logger.Debug("Postgres was inited")

	if err = migratePostgres(db); err != nil {
		logger.Fatal("Migrations error", err)
	}

	repos := repository.NewRepository(db)
	logger.Debug("Repositories was inited")

	services := service.NewServices(repos, cfg.Salt, cfg.SignKey, cfg.TokenTTL)
	logger.Debug("Services was inited")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := v1.NewHandler(ctx, services, logger.Logger).InitRoutes()
	logger.Debug("Handlers was inited")

	httpServer := http_server.New(router, cfg.Http.Port)
	logger.Info(fmt.Sprintf("Http server was started at port: %s", cfg.Http.Port))

	notify := make(chan os.Signal, 1)
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)

	select {
	case <-notify:
		logger.Debug("Stop signal received")
	case <-httpServer.Notify():
		logger.Error("Server running error", slog.Any("msg", err.Error()))
	}

	err = httpServer.Shutdown()
	if err != nil {
		logger.Error("Error while shutting down the server", slog.Any("msg", err.Error()))
	}
}

func connectPostgres(cfg config.Postgres) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DbName, cfg.Username, cfg.Password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func migratePostgres(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations/migrate", "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
