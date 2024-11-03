package handler

import (
	stools "github.com/antalkon/zic_server/internal/services/sTools"
	"github.com/gin-gonic/gin"
)

func Clear(c *gin.Context) {
	stools.Clear(c)
}
