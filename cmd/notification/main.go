package main

import (
	"notification-service/internal/config"
	"notification-service/internal/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env, cfg.LogPath)

	log.Info("starting notification service")
}
