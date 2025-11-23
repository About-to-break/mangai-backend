package main

import (
	"backend/internal/config"
	"backend/internal/logger"
	"log/slog"
)

func main() {
	cfg := config.LoadConfig()
	logger.SetupLogger(cfg.LogLevel)
	slog.Info("Startup complete")
}
