package scloud

import (
	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func TasksCount(c *gin.Context) {
	all, pend, err := cloudpg.GetTasksCount()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"all": all, "pending": pend, "color": "green-500", "success": true})
}
