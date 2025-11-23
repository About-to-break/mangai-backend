package mangalocal

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

// UploadView обрабатывает загрузку изображения
// @Summary Upload an image
// @Description Uploads an image file to the server
// @Tags LocalManga
// @Accept multipart/form-data
// @Param file formData file true "Image file"
// @Produce json
// @Success 200 {object} map[string]string "File uploaded"
// @Failure 400 {object} map[string]string "No file received"
// @Router /local/upload [post]
func UploadView(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "uploaded",
		"filename": file.Filename,
	})

}
