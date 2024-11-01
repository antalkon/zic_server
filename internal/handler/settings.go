package handler

import (
	ssettings "github.com/antalkon/zic_server/internal/services/sSettings"
	"github.com/gin-gonic/gin"
)

// post статуса системы
func SysStaus(c *gin.Context) {
	ssettings.SystemStatus(c)
}

// страница настроек
func SettingsPage(c *gin.Context) {
	ssettings.SettingsPage(c)
}

// post данных тг
func TgData(c *gin.Context) {
	ssettings.TgData(c)
}

// put обновления данных тг
func TgNewData(c *gin.Context) {
	ssettings.TgNewData(c)
}

// post настроек сети
func SecNetwork(c *gin.Context) {
	ssettings.Network(c)
}

// post настроек безопасности
func SecFirewall(c *gin.Context) {
	ssettings.Firewall(c)
}

// post настроек авторизации
func SecAuth(c *gin.Context) {
	ssettings.Auth(c)
}

// post настрок безопасности гет
func SecGet(c *gin.Context) {
	ssettings.Get(c)
}

// получение версии
func InfoVesion(c *gin.Context) {
	ssettings.Version(c)
}
