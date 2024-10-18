package handler

import (
	sdashboard "github.com/antalkon/zic_server/internal/services/sDashboard"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	sdashboard.Dashboard(c)
}
