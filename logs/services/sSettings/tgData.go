package ssettings

import (
	"net/http"

	"github.com/antalkon/zic_server/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// TgData возвращает конфигурацию Telegram бота без полей name и password
func TgData(c *gin.Context) {
	// Загружаем конфигурацию
	err := config.LoadDataConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	// Извлекаем конфигурацию Telegram бота напрямую
	token := viper.GetString("tg_bot.token")
	sysSt := viper.GetBool("tg_bot.sysStat")
	newTask := viper.GetBool("tg_bot.newtask")

	// Формируем ответ с конфигурацией бота
	response := gin.H{
		"success": true,
		"token":   token,
		"sysSt":   sysSt,
		"newTask": newTask,
	}

	// Возвращаем JSON-ответ с конфигурацией бота
	c.JSON(http.StatusOK, response)
}
