package spc

import (
	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
)

func Del(c *gin.Context) {
	id := c.Param("id")
	err := pcpg.DeletePcById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "PC deleted successfully"})
}
