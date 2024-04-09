package app

import (
	"io"
	"log"
	"log/slog"
	"sso/internal/lib/logger/services"
)

func InitLogger(
	LevelDebug slog.Leveler,
	OutputWriter io.Writer,
) *slog.Logger {
	return slog.New(&services.PrettyLogger{
		Handler: slog.NewJSONHandler(OutputWriter, &slog.HandlerOptions{Level: LevelDebug}),
		Logger:  log.New(OutputWriter, "", 0),
	})
}
