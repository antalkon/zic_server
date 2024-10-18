package sdashboard

import "github.com/gin-gonic/gin"

func Dashboard(c *gin.Context) {
	c.HTML(200, "dashboard.html", gin.H{
		"title":   "Панель управления",
		"setting": false,
	})
}
