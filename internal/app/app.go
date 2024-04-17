package app

import (
	"log/slog"
	psqlapp "sso/internal/app/database/psql"
	redisapp "sso/internal/app/database/redis"
	grpcapp "sso/internal/app/grpc"
	"sso/internal/config"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
	PSQLApp    *psqlapp.PsqlApp
	REDISApp   *redisapp.RedisApp
}

func NewApp(
	log *slog.Logger,
	gRPCPort int,
	PsqlConfig config.PSQL,
	RedisConfig config.Redis,
	tokenTtl time.Duration,
) *App {
	psqlApp := psqlapp.NewPsqlApp(
		log,
		PsqlConfig.Username,
		PsqlConfig.Dbname,
		PsqlConfig.Password,
		PsqlConfig.Host,
		PsqlConfig.Port,
	)

	redisApp := redisapp.NewRedisApp(
		log,
		RedisConfig.Addr,
		RedisConfig.Password,
		RedisConfig.Db,
	)

	grpcApp := grpcapp.NewGrpcApp(
		gRPCPort,
		config.NewHandleServiceStructure(
			psqlApp,
			redisApp,
			tokenTtl,
			log,
		),
	)

	return &App{GRPCServer: grpcApp, PSQLApp: psqlApp, REDISApp: redisApp}
}
