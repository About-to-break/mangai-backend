package logger

import (
	"log/slog"
	"os"
	"strings"
)

var Logger *slog.Logger

func SetupLogger(level string) {
	var lvl slog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     lvl,
	})

	Logger = slog.New(handler)

	slog.SetDefault(Logger)

	slog.Info("Logger initialized", "level", level)
}
