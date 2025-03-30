package scloud

import (
	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	cls, err := cloudpg.GetAllClass()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, cls)
}
