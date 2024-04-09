package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"sso/internal/lib/logger/config"
	"time"
)

type PrettyLogger struct {
	Handler slog.Handler
	Logger  *log.Logger
	Attrs   []slog.Attr
}

func (p PrettyLogger) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (p PrettyLogger) Handle(ctx context.Context, record slog.Record) error {
	data, err := p.formatData(record)

	if err != nil {
		return err
	}

	p.Logger.Println(
		p.fillColor(time.Now().Format("2006-01-02 15:04:05.000"), config.LevelForColor[record.Level]),
		p.fillColor(record.Level.String()+":", config.LevelForColor[record.Level]),
		p.fillColor(record.Message, config.LevelForColor[record.Level]),
		p.fillColor(string(data), config.White),
	)

	return nil
}

func (p PrettyLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyLogger{
		Handler: p.Handler,
		Logger:  p.Logger,
		Attrs:   attrs,
	}
}

func (p PrettyLogger) WithGroup(name string) slog.Handler {
	return &PrettyLogger{
		Handler: p.Handler.WithGroup(name),
		Logger:  p.Logger,
	}
}

func (p PrettyLogger) fillColor(text string, color int) string {
	return fmt.Sprintf(
		"\033[%dm%s\033[0m",
		color,
		text,
	)
}

func (p PrettyLogger) formatData(record slog.Record) ([]byte, error) {
	if record.NumAttrs() <= 0 {
		return make([]byte, 0), nil
	}

	fields := make(map[string]interface{}, record.NumAttrs())

	record.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range p.Attrs {
		fields[a.Key] = a.Value.Any()
	}

	b, err := json.MarshalIndent(fields, "", "  ")

	if err != nil {
		return nil, err
	}

	return b, nil
}
