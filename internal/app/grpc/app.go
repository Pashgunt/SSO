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

//✔ Создаёт gRPC-сервер
//✔ Регистрирует сервисы
//✔ Запускает и останавливает сервер
//✔ Логирует события

type HandlerServices interface {
	MakeAuthService() *authservice.Auth
}

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

//Создаётся новый grpc.Server().
//Инициализируется структура, управляющая сервисами (HandlerServices).
//Регистрируется сервис аутентификации (RegisterServerApiHandler).
//Возвращается объект App.

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

//Создаётся лог с названием текущего метода.
//Открывается TCP-соединение (net.Listen) на заданном порту.
//Запускается сервер (a.gRPCServer.Serve(listener)).
//Если что-то идёт не так — ошибка передаётся наверх.

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
