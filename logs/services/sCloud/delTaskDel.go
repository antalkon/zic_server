package scloud

import (
	"strconv"

	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func DelTaskDel(c *gin.Context) {
	id := c.Param("id")
	myUInt64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	err = cloudpg.DelTaskDel(myUInt64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Задание успешно удалено"})
}
