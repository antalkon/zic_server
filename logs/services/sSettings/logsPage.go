package ssettings

import "github.com/gin-gonic/gin"

func SettingsLogsPage(c *gin.Context) {

	c.HTML(200, "settings_logs.html", gin.H{
		"title":   "Логи панели управления",
		"page":    "settings_logs",
		"setting": true,
	})
}
