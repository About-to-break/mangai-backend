package mangalocal

import (
	"backend/internal/config"
	"backend/internal/services"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func SetupLocalMangaRoutes(routerGroup *gin.RouterGroup, cfg *config.Config) {

	uploadStorage, err := services.NewMinioStorage(
		cfg.MinioEndpoint,
		cfg.MinioAccessKey,
		cfg.MinioSecretKey,
		cfg.MinioUseSSL,
	)
	if err != nil {
		slog.Error("Setting up storage failed", err)
	}

	uploadQueue, err := services.NewRabbitMQueue(
		cfg.RabbitURI,
	)
	if err != nil {
		slog.Error("Setting up queue failed", err)
	}

	imageUploadCtrl := UploadController{
		Storage:  uploadStorage,
		Queue:    uploadQueue,
		Bucket:   cfg.MinioBucket,
		Exchange: cfg.RabbitExchange,
		Key:      cfg.RabbitRoutingKey,
	}
	localMangaGroup := routerGroup.Group("local")
	{
		localMangaGroup.GET("/", IndexView)
		localMangaGroup.POST("/upload", imageUploadCtrl.UploadView)

	}
}
