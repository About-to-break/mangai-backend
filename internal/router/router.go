package router

import (
	"github.com/gin-gonic/gin"

	"backend/internal/apps/mangalocal"

	_ "backend/docs" // пакет с сгенерированным swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	// Routes from apps
	appRoutes := router.Group("/")
	{
		// Here go app routers groups
		mangalocal.LocalMangaRoutes(appRoutes)
	}
	// Extra routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
