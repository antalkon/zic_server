package scloud

import (
	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func TaskClass(c *gin.Context) {
	id := c.Param("id")
	tasks, err := cloudpg.GetClassTasks(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)

}
