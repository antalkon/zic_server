package spc

import (
	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/antalkon/zic_server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddPc(c *gin.Context) {
	var pc models.Computer

	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := utils.GenerateRandomNumber(12)
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to generate random ID"})
	}
	pc.ID = uint64(id)
	if err := pcpg.AddNewPc(&pc); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "Pc added successfully"})
}
