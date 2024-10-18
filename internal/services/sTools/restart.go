package stools

import (
	"github.com/antalkon/zic_server/pkg/tools"
	"github.com/gin-gonic/gin"
)

func Restart(c *gin.Context) {
	err := tools.Restart()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "Restart successful"})
}
