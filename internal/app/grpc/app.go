package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"runtime"
	"sso/internal/config"
	"sso/internal/grpc/auth"
	"sso/internal/services"
	authservice "sso/internal/services/auth"
)

type HandlerServices interface {
	MakeAuthService() *authservice.Auth
}

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewGrpcApp(
	port int,
	handleServiceStructure *config.HandleServiceStructure,
) *App {
	gRPCServer := grpc.NewServer()
	handlerServices := services.NewHandlerServices(handleServiceStructure)

	auth.RegisterServerApiHandler(gRPCServer, handlerServices.MakeAuthService())

	return &App{
		log:        handleServiceStructure.Log,
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
	log := a.log.With(slog.String("operation", a.MethodForLog()))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		return fmt.Errorf("%s: %w", a.MethodForLog(), err)
	}

	log.Info(
		"Running gRPC server",
		slog.Int("PORT", a.port),
		slog.String("addr", listener.Addr().String()),
	)

	if err := a.gRPCServer.Serve(listener); err != nil {
		return fmt.Errorf("%s: %w", a.MethodForLog(), err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.With(slog.String("operation", a.MethodForLog())).
		Info(
			"Stopping gRPC server",
			slog.Int("PORT", a.port),
		)

	a.gRPCServer.GracefulStop()
}

func (a *App) MethodForLog() string {
	pc, _, _, _ := runtime.Caller(1)

	return runtime.FuncForPC(pc).Name()
}
