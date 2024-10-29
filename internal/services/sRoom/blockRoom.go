package sroom

import (
	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
)

func RoomBlock(c *gin.Context) {
	id := c.Param("id")

	// Заблокировать все ПК с room_id = id
	err := roompg.ChangePCRoomStatus(id, "blocked")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ПК в комнате заблокированы"})
}

func RoomUnBlock(c *gin.Context) {
	id := c.Param("id")

	// Заблокировать все ПК с room_id = id
	err := roompg.ChangePCRoomStatus(id, "work")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ПК в комнате заблокированы"})
}
