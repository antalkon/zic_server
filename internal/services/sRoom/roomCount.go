package sroom

import (
	"fmt"

	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func RoomCount(c *gin.Context) {
	var room models.RoomCount

	err, count, on := roompg.GetRoom()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Errorf("Не удалось обновить информацию активных Rooms. %w", err.Error())})
		return
	}
	if on == 0 {
		room.Color = "red-500"
	} else {
		room.Color = "green-500"

	}
	room.Count = count
	room.On = on

	c.JSON(200, gin.H{
		"success": true, // Поле "success" указывает на успешное выполнение
		"count":   room.Count,
		"on":      room.On,
		"color":   room.Color,
	})
}
