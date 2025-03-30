package sroom

import (
	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
)

func RoomData(c *gin.Context) {
	rooms, err := roompg.GetRooms()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, rooms)
}
