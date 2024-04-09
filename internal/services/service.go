package services

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	psqlapp "sso/internal/app/database/psql"
	redisapp "sso/internal/app/database/redis"
	authservice "sso/internal/services/auth"
	"time"
)

type HandlerServices struct {
	psql     *psqlapp.PsqlApp
	redisApp *redisapp.RedisApp
	tokenTtl time.Duration
	log      *slog.Logger
}

func NewHandlerServices(psql *psqlapp.PsqlApp, redisApp *redisapp.RedisApp, tokenTtl time.Duration, log *slog.Logger) *HandlerServices {
	return &HandlerServices{psql: psql, redisApp: redisApp, tokenTtl: tokenTtl, log: log}
}

func (hs *HandlerServices) MakeAuthService() *authservice.Auth {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: hs.psql.Db(),
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &authservice.Auth{
		Gorm:     gormDB,
		Redis:    hs.redisApp.Client(),
		TokenTtl: hs.tokenTtl,
		Log:      hs.log,
	}
}
