package ssettings

import "github.com/gin-gonic/gin"

func SettingsPage(c *gin.Context) {
	c.HTML(200, "settings.html", gin.H{
		"title":   "Панель управления",
		"setting": false,
	})
}
