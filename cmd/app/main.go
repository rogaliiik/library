package main

import (
	"flag"

	"github.com/rogaliiik/library/internal/app"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config")
	flag.Parse()
}

func main() {
	app.Run(configPath)
}
