package grpcapp

import (
	"fmt"
	"net"
	notificationGRPC "notification-service/internal/grpc/notification"
	templateGRPC "notification-service/internal/grpc/template"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	log        *zap.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.Logger, port int) *App {

	grpcServer := grpc.NewServer()

	notificationGRPC.Register(grpcServer)
	templateGRPC.Register(grpcServer)

	return &App{
		log:        log,
		port:       port,
		gRPCServer: grpcServer,
	}
}

func (a *App) MustStart() {
	const op = "grpcapp.MustStart"

	if err := a.Start(); err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}
}

func (a *App) Start() error {
	const op = "grpcapp.Start"

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	a.log.Info("grpc server started", zap.String("address", listen.Addr().String()))

	if err := a.gRPCServer.Serve(listen); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
	a.log.Info("grpc server stopped")
}
