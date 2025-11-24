package main

import (
	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/router"
	"log/slog"
)

func main() {
	cfg := config.LoadConfig()

	logger.SetupLogger(cfg.LogLevel)

	r := router.SetupRouters(cfg)

	err := r.Run(":" + cfg.ServerPort)

	slog.Info("Startup complete")

	if err != nil {
		return
	}
}
