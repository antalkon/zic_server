package sload

import (
	loadpg "github.com/antalkon/zic_server/internal/db/db_pg/load_pg"
	"github.com/gin-gonic/gin"
)

func RamLoad(c *gin.Context) {
	ram, err := loadpg.RamLoad24()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, ram)
}
