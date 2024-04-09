package config

import "log/slog"

const (
	Rest = iota + 30
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)

var LevelForColor = map[slog.Level]int{
	slog.LevelDebug: Blue,
	slog.LevelInfo:  Cyan,
	slog.LevelWarn:  Yellow,
	slog.LevelError: Red,
}
