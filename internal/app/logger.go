package app

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"time"
)

const (
	prod = "prod"
	dev  = "dev"

	logsDir     = "logs"
	logTemplate = "library-%s.log"
)

type Logger struct {
	*slog.Logger
}

func (l *Logger) Fatal(msg string, err error, args ...any) {
	args = append(args, slog.Any("error", err.Error()))
	l.Error(msg, args...)
	os.Exit(1)
}

func NewLogger(logLevel string) (*Logger, error) {
	var logger *slog.Logger
	switch logLevel {
	case prod:
		f, err := createLogFile()
		if err != nil {
			return nil, err
		}
		logger = slog.New(
			slog.NewJSONHandler(io.MultiWriter(f, os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case dev:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return &Logger{logger}, nil
}

func createLogFile() (*os.File, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path.Join()
	err = os.MkdirAll(path.Join(curDir, logsDir), os.ModePerm)
	if err != nil {
		return nil, err
	}

	return os.Create(fmt.Sprintf(path.Join(logsDir, logTemplate), time.Now().Format("20060102-1504")))
}
