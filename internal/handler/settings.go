package handler

import (
	ssettings "github.com/antalkon/zic_server/internal/services/sSettings"
	"github.com/gin-gonic/gin"
)

func SysStaus(c *gin.Context) {
	ssettings.SystemStatus(c)
}
