package main

import (
	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/log"
	defaultLogger "log"
)

func main() {
	cfg := getConfig()
	logger := getLogger(cfg)
	app.Start(cfg, logger)
}

func getLogger(cfg config.Config) log.Logger {
	loggingLevel, err := log.StringToLoggingLevel(cfg.Logging.Level)
	if err != nil {
		defaultLogger.Fatalf("unknown logging level: %v", err)
	}
	logger := log.NewLogger()
	logger.SetLevel(loggingLevel)
	return logger
}

func getConfig() config.Config {
	cfg, err := config.Load()
	if err != nil {
		defaultLogger.Fatalf("config load error: %v", err)
	}
	return cfg
}
