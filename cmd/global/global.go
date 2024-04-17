package global

import (
	"log/slog"
	"sso/internal/config"
)

func InitGlobalStructures() (*config.Config, *slog.Logger) {
	cfg := config.MustLoadConfig()
	logger := config.MustLoadLogger(cfg.Env)

	return cfg, logger
}
