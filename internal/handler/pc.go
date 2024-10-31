package handler

import (
	spc "github.com/antalkon/zic_server/internal/services/sPc"
	"github.com/gin-gonic/gin"
)

func AddNewPc(c *gin.Context) {
	spc.AddPc(c)
}

func PcCount(c *gin.Context) {
	spc.PcCount(c)
}

func ServerPing(c *gin.Context) {
	spc.PcPing(c)
}

func PcOff(c *gin.Context) {
	spc.OffPc(c)
}

func PcReboot(c *gin.Context) {
	spc.RebootPc(c)
}
