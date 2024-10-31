package spc

import (
	"fmt"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
)

func StartPc(c *gin.Context) {
	lip := c.Param("lip")

	d, err := pcpg.GetPcData(lip)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err, d)
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"status":  d.Status,
	})
}
