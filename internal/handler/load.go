package handler

import (
	sload "github.com/antalkon/zic_server/internal/services/sLoad"
	"github.com/gin-gonic/gin"
)

func LoadCPU(c *gin.Context) {
	sload.CpuLoad(c)
}
func LoadRAM(c *gin.Context) {
	sload.RamLoad(c)
}

func LoadNetwork(c *gin.Context) {
	sload.NetworkLoad(c)
}

func LoadStorage(c *gin.Context) {
	sload.StorageLoad(c)
}
