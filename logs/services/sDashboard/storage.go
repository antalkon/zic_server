package sdashboard

import "github.com/gin-gonic/gin"

func LoadStoragePage(c *gin.Context) {
	c.HTML(200, "storage_load.html", gin.H{
		"title":   "Загрузка хранилища",
		"setting": true,
	})
}
