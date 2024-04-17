package config

import (
	"log/slog"
	psqlapp "sso/internal/app/database/psql"
	redisapp "sso/internal/app/database/redis"
	"time"
)

type HandleServiceStructure struct {
	Psql     *psqlapp.PsqlApp
	RedisApp *redisapp.RedisApp
	TokenTtl time.Duration
	Log      *slog.Logger
}

func NewHandleServiceStructure(
	psql *psqlapp.PsqlApp,
	redisApp *redisapp.RedisApp,
	tokenTtl time.Duration,
	log *slog.Logger,
) *HandleServiceStructure {
	return &HandleServiceStructure{
		Psql:     psql,
		RedisApp: redisApp,
		TokenTtl: tokenTtl,
		Log:      log,
	}
}
