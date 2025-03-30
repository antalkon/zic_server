package sroom

import (
	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
)

func DelRoom(c *gin.Context) {
	id := c.Param("id")

	if err := roompg.DelRoomById(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "Комната с ID " + id + " удалена"})
}
