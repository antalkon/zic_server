package ssettings

import (
	"net/http"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Auth(c *gin.Context) {
	var g models.ReqAuth
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.LoadDataConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	viper.Set("sec.reqauth", g.Auth)

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
