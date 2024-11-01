package handler

import (
	sload "github.com/antalkon/zic_server/internal/services/sLoad"
	"github.com/gin-gonic/gin"
)

// *** ЗА 24 ЧАСА ***

// post нагрузки cpu
func LoadCPU(c *gin.Context) {
	sload.CpuLoad(c)
}

// post нагрузки ram
func LoadRAM(c *gin.Context) {
	sload.RamLoad(c)
}

// post нагрузки сети
func LoadNetwork(c *gin.Context) {
	sload.NetworkLoad(c)
}

// post нагрузки диска
func LoadStorage(c *gin.Context) {
	sload.StorageLoad(c)
}
