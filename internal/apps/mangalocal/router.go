package mangalocal

import (
	"github.com/gin-gonic/gin"
)

func LocalMangaRoutes(routerGroup *gin.RouterGroup) {
	localMangaGroup := routerGroup.Group("local")
	{
		localMangaGroup.GET("/", IndexView)
		localMangaGroup.POST("/upload", UploadView)

	}
}
