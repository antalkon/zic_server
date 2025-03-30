package sroom

import (
	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func NewRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := utils.GenerateRandomNumber(12)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate random ID"})
	}
	room.ID = uint64(id)
	if err := roompg.NewRoom(&room); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Комната добавлена успешно!"})
}
