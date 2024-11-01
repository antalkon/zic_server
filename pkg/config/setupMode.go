// pkg/config/config.go

package config

import (
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// SetupGinMode устанавливает режим работы Gin на основе конфигурации
func SetupGinMode(log logger.Logger) {
	mode := viper.GetString("http_server.mode")
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
		log.Info("Режим сервера: Разработка (Debug)")
	case "prod":
		gin.SetMode(gin.ReleaseMode)
		log.Info("Режим сервера: Продакшен (Release)")
	case "test":
		gin.SetMode(gin.TestMode)
		log.Info("Режим сервера: Тестирование (Test)")
	default:
		gin.SetMode(gin.ReleaseMode)
		log.Warn("Неопознанный режим. Установлен режим по умолчанию: Продакшен (Release)")
	}
}
