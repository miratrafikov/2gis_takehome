package main

import (
	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/config"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := getConfig()
	app.Start(ctx, cfg)
}

func getConfig() config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}
	return cfg
}
