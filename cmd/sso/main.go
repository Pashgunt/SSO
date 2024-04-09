package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)

func main() {
	cfg := config.MustLoadConfig()
	logger := config.MustLoadLogger(cfg.Env)
	application := app.NewApp(
		logger,
		cfg.GRPC.Port,
		cfg.PSQL,
		cfg.Redis,
		cfg.TokenTtl,
	)

	defer application.PSQLApp.Stop()

	go application.PSQLApp.MustRun()
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-stop:
		application.GRPCServer.Stop()
		application.PSQLApp.Stop()
		application.REDISApp.Stop()
		logger.Info("Application was stopped", slog.String("signal", s.String()))
	}
}
