package mangalocal

import (
	"backend/internal/config"
	"backend/internal/services"
	"github.com/gin-gonic/gin"
)

func LocalMangaRoutes(routerGroup *gin.RouterGroup, cfg *config.Config, storage services.StorageService) {
	imageUploadCtrl := UploadController{
		Storage: storage,
		Bucket:  cfg.MinioBucket,
	}
	localMangaGroup := routerGroup.Group("local")
	{
		localMangaGroup.GET("/", IndexView)
		localMangaGroup.POST("/upload", imageUploadCtrl.UploadView)

	}
}
