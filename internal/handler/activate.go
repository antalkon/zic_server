package handler

import (
	sactivate "github.com/antalkon/zic_server/internal/services/sActivate"
	"github.com/gin-gonic/gin"
)

func Activate(c *gin.Context) {
	sactivate.Activation(c)
}

func ActivatePage(c *gin.Context) {
	sactivate.ActPage(c)
}
