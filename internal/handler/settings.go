package handler

import (
	ssettings "github.com/antalkon/zic_server/internal/services/sSettings"
	"github.com/gin-gonic/gin"
)

func SysStaus(c *gin.Context) {
	ssettings.SystemStatus(c)
}

func SettingsPage(c *gin.Context) {
	ssettings.SettingsPage(c)
}

func TgData(c *gin.Context) {
	ssettings.TgData(c)
}

func TgNewData(c *gin.Context) {
	ssettings.TgNewData(c)
}
