package scloud

import (
	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/gin-gonic/gin"
)

func ClassPage(c *gin.Context) {
	classes, err := cloudpg.GetAllClass()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "c_class.html", gin.H{
		"title": "Классы",
		"class": classes,
	})
}
