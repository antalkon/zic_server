package handler

import (
	stools "github.com/antalkon/zic_server/internal/services/sTools"
	"github.com/gin-gonic/gin"
)

// post перезагрузки
func Restart(c *gin.Context) {
	stools.Restart(c)
}
