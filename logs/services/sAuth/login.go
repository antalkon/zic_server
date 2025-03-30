package sauth

import (
	"log"
	"net/http"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/antalkon/zic_server/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Login(c *gin.Context) {
	var l models.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.LoadDataConfig(); err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load data config"})
		return
	}
	// Получаем значения полей из файла конфигурации
	name := viper.GetString("name")
	password := viper.GetString("password")
	if name != l.Name || password != l.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := sessions.SetUserSession(c, l.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})

}
