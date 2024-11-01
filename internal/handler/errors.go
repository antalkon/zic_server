package handler

import (
	serrors "github.com/antalkon/zic_server/internal/services/sErrors"
	"github.com/gin-gonic/gin"
)

// Страница ошибки отказано в доступе
func NotAuth(c *gin.Context) {
	serrors.Page403(c)
}

// Страница ошибки не найдено
func NotFound(c *gin.Context) {
	serrors.Page404(c)
}
