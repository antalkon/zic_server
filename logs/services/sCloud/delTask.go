package scloud

import (
	"strconv"

	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func DelTask(c *gin.Context) {
	id := c.Param("id")
	classID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := cloudpg.DelTask(classID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}
