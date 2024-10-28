package sdashboard

import "github.com/gin-gonic/gin"

func LoadRAMPage(c *gin.Context) {
	c.HTML(200, "ram_load.html", gin.H{
		"title":   "Загрузка оперативной памяти",
		"setting": true,
	})
}
