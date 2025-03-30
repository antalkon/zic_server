package spc

import (
	getip "github.com/antalkon/zic_server/pkg/getIP"
	"github.com/gin-gonic/gin"
)

func PcPing(c *gin.Context) {
	pip, lip, err := getip.GetIps()
	_ = pip
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Pong",
		"lip":     lip,
		"port":    "8385",
	})
}
