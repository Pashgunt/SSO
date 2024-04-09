package redisapp

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisApp struct {
	log    *slog.Logger
	client *redis.Client
}

func (redis *RedisApp) Client() *redis.Client {
	return redis.client
}

func NewRedisApp(
	log *slog.Logger,
	addr string,
	password string,
	db int,
) *RedisApp {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	log.With(slog.String("operation", "psqlapp.NewRedisApp")).
		Info(
			"Starting REDIS server",
			slog.String("conn", addr),
		)

	return &RedisApp{log: log, client: client}
}

func (redis *RedisApp) Stop() {
	redis.log.With(slog.String("operation", "redisapp.Stop")).
		Info(
			"Stopping REDIS server",
			slog.String("conn", redis.client.String()),
		)

	if err := redis.client.Close(); err != nil {
		panic(err)
	}
}
