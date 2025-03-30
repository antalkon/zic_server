package spc

import (
	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
)

func PcPage(c *gin.Context) {
	id := c.Param("id")
	pc, err := pcpg.GetPcDataId(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.HTML(200, "pc.html", gin.H{
		"data":  pc,
		"title": "PC Page",
	})
}
