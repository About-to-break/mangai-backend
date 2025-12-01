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

	slog.Info("Startup complete")

	err := r.Run("0.0.0.0:" + cfg.ServerPort)

	if err != nil {
		return
	}

}
