package spc

import (
	"fmt"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func PcCount(c *gin.Context) {
	var pc models.PcCount

	err, count, on := pcpg.GetPcCountAndOn()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Errorf("Не удалось обновить информацию активных ПК. %w", err.Error())})
		return
	}
	if on == 0 {
		pc.Color = "red-500"
	} else {
		pc.Color = "green-500"

	}
	pc.Count = count
	pc.On = on

	c.JSON(200, gin.H{
		"success": true, // Поле "success" указывает на успешное выполнение
		"count":   pc.Count,
		"on":      pc.On,
		"color":   pc.Color,
	})
}
