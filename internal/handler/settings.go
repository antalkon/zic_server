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

func SecNetwork(c *gin.Context) {
	ssettings.Network(c)
}

func SecFirewall(c *gin.Context) {
	ssettings.Firewall(c)
}

func SecAuth(c *gin.Context) {
	ssettings.Auth(c)
}

func SecGet(c *gin.Context) {
	ssettings.Get(c)
}

func InfoVesion(c *gin.Context) {
	ssettings.Version(c)
}
