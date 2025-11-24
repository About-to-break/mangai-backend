package main

import (
	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/router"
	"backend/internal/services"
	"log/slog"
)

func main() {
	cfg := config.LoadConfig()

	logger.SetupLogger(cfg.LogLevel)

	storage, _ := services.NewMinioStorage(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioUseSSL,
	)

	r := router.SetupRouters(cfg, storage)

	err := r.Run(":" + cfg.ServerPort)

	slog.Info("Startup complete")

	if err != nil {
		return
	}
}
