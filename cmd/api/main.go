package main

import (
	"flag"
	"github.com/oli4maes/winedata/internal/jsonlog"
	"os"
	"sync"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	wg     sync.WaitGroup
	logger *jsonlog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
