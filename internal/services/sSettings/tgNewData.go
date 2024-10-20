package ssettings

import (
	"net/http"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TgNewData(c *gin.Context) {
	var d models.NewDataTg
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Загружаем конфигурацию
	err := config.LoadDataConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	// Устанавливаем новые значения
	viper.Set("tg_bot.token", d.Token)
	viper.Set("tg_bot.sysSt", d.SysStat)
	viper.Set("tg_bot.newTask", d.NewTasks)

	// Сохраняем изменения в файл
	if err := config.SaveDataConfig(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save configuration"})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(200, gin.H{
		"success": true,
		"info":    "Успешно!",
	})
}
