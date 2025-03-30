package spc

import (
	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	var pc models.Computer
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := pcpg.UpdatePcById(pc.ID, &pc)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "PC updated successfully"})
}
