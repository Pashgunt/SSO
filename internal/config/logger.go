package config

import (
	"log/slog"
	"os"
	"sso/internal/lib/logger/app"
)

const (
	envLocal   = "local"
	envDev     = "dev"
	envPreprod = "preprod"
	envProd    = "prod"
)

func MustLoadLogger(envCfg string) *slog.Logger {
	switch envCfg {
	case envLocal:
		return app.InitLogger(slog.LevelDebug, os.Stdout)
	case envDev:
		return app.InitLogger(slog.LevelDebug, os.Stdout)
	case envPreprod:
		return app.InitLogger(slog.LevelWarn, os.Stdout)
	case envProd:
		return app.InitLogger(slog.LevelError, os.Stdout)
	default:
		return app.InitLogger(slog.LevelDebug, os.Stdout)
	}
}
