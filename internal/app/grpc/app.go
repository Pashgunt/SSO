package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sso/internal/grpc/auth"
	authservice "sso/internal/services/auth"
)

type HandlerServices interface {
}

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGrpcApp(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()
	authService := &authservice.Auth{}

	auth.RegisterServerApiHandler(gRPCServer, authService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	log := a.log.With(slog.String("operation", "grpcapp.Run"))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	log.Info(
		"Running gRPC server",
		slog.Int("PORT", a.port),
		slog.String("addr", listener.Addr().String()),
	)

	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", "grpcapp.Run", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With(slog.String("operation", "grpcapp.Stop")).
		Info(
			"Stopping gRPC server",
			slog.Int("PORT", a.port),
		)

	a.gRPCServer.GracefulStop()
}
