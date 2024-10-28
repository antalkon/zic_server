package sload

import (
	"net/http"

	loadpg "github.com/antalkon/zic_server/internal/db/db_pg/load_pg"
	"github.com/gin-gonic/gin"
)

func NetworkLoad(c *gin.Context) {
	networkData, err := loadpg.NetworkLoad24()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, networkData)
}
