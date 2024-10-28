package sdashboard

import "github.com/gin-gonic/gin"

func LoadNetworkPage(c *gin.Context) {
	c.HTML(200, "network_load.html", gin.H{
		"title":   "Загрузка сети",
		"setting": true,
	})
}
