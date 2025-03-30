package app

import (
	grpcapp "notification-service/internal/app/grpc"
	"notification-service/internal/config"

	"go.uber.org/zap"
)

type App struct {
	GRPCApp *grpcapp.App
}

func New(log *zap.Logger, cfg *config.Config) *App {
	grpcApp := grpcapp.New(log, cfg.GRPC.Port)

	return &App{
		GRPCApp: grpcApp,
	}
}
