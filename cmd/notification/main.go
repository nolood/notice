package main

import (
	"notification-service/internal/app"
	"notification-service/internal/config"
	"notification-service/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env, cfg.LogPath)

	application := app.New(log, cfg)

	application.GRPCApp.MustStart()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	application.GRPCApp.Stop()
}
