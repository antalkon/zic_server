package handler

import (
	sdashboard "github.com/antalkon/zic_server/internal/services/sDashboard"
	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	sdashboard.Dashboard(c)
}

func LoadCPUPage(c *gin.Context) {
	sdashboard.LoadCPUPage(c)
}

func LoadRAMPage(c *gin.Context) {
	sdashboard.LoadRAMPage(c)
}

func LoadNetworkPage(c *gin.Context) {
	sdashboard.LoadNetworkPage(c)
}

func LoadStoragePage(c *gin.Context) {
	sdashboard.LoadStoragePage(c)
}
