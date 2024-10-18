package spc

import (
	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func PcCount(c *gin.Context) {
	var pc models.PcCount

	err, count, on := pcpg.GetPcCountAndOn()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if count == 0 {
		pc.Color = "bg-red-500"
	} else {
		pc.Color = "bg-green-500"

	}
	pc.Count = count
	pc.On = on

	c.JSON(200, pc)

}
