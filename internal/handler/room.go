package handler

import (
	sroom "github.com/antalkon/zic_server/internal/services/sRoom"
	"github.com/gin-gonic/gin"
)

func AddNewRoom(c *gin.Context) {
	sroom.NewRoom(c)
}

func RoomCount(c *gin.Context) {
	sroom.RoomCount(c)
}

func RoomsData(c *gin.Context) {
	sroom.RoomData(c)
}

func RoomDataPage(c *gin.Context) {
	sroom.RoomPage(c)
}

func BlockRoom(c *gin.Context) {
	sroom.RoomBlock(c)
}

func UnblockRoom(c *gin.Context) {
	sroom.RoomUnBlock(c)
}
