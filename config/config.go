package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Http
	Postgres
	Log
	JWT
}

type Http struct {
	Port string `yaml:"port"`
}

type Postgres struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Log struct {
	Level string `yaml:"level"`
}

type JWT struct {
	Salt     string        `yaml:"salt"`
	SignKey  string        `yaml:"sign_key"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
