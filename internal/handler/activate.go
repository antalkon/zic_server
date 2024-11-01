package handler

import (
	sactivate "github.com/antalkon/zic_server/internal/services/sActivate"
	"github.com/gin-gonic/gin"
)

// Активируем сервер
func Activate(c *gin.Context) {
	sactivate.Activation(c)
}

// Страница активации сервера
func ActivatePage(c *gin.Context) {
	sactivate.ActPage(c)
}
