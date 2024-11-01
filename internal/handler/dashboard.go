package handler

import (
	sdashboard "github.com/antalkon/zic_server/internal/services/sDashboard"
	sdata "github.com/antalkon/zic_server/internal/services/sData"
	"github.com/gin-gonic/gin"
)

// Страница дашборда
func Dashboard(c *gin.Context) {
	sdashboard.Dashboard(c)
}

// Страница статистики cpu
func LoadCPUPage(c *gin.Context) {
	sdashboard.LoadCPUPage(c)
}

// Страница статистики ram
func LoadRAMPage(c *gin.Context) {
	sdashboard.LoadRAMPage(c)
}

// Страница статистки сети
func LoadNetworkPage(c *gin.Context) {
	sdashboard.LoadNetworkPage(c)
}

// Страница статистики хранилища
func LoadStoragePage(c *gin.Context) {
	sdashboard.LoadStoragePage(c)
}

// Страница данных
func DataPage(c *gin.Context) {
	sdata.LoadDataPage(c)
}
