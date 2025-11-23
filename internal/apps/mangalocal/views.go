package mangalocal

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

// IndexView отдаёт статический frontend
// @Summary Отдать главную страницу локального приложения
// @Description Возвращает index.html
// @Tags LocalManga
// @Produce html
// @Success 200 {string} string "index.html"
// @Router /local/ [get]
func IndexView(c *gin.Context) {
	c.File(filepath.Join("frontend", "index.html"))
}
