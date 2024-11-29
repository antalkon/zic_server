package handler

import (
	sroom "github.com/antalkon/zic_server/internal/services/sRoom"
	"github.com/gin-gonic/gin"
)

// post создания новой комнаты
func AddNewRoom(c *gin.Context) {
	sroom.NewRoom(c)
}

// post количества компант
func RoomCount(c *gin.Context) {
	sroom.RoomCount(c)
}

// post получения всех комнат
func RoomsData(c *gin.Context) {
	sroom.RoomData(c)
}

// страница комнаты
func RoomDataPage(c *gin.Context) {
	sroom.RoomPage(c)
}

// post блокировки комнаты
func BlockRoom(c *gin.Context) {
	sroom.RoomBlock(c)
}

// post разблокировки комнаты
func UnblockRoom(c *gin.Context) {
	sroom.RoomUnBlock(c)
}

// post выключения комнаты
func OffRoom(c *gin.Context) {
	sroom.RoomOff(c)
}

// post перезагрузки комнаты
func RebootRoom(c *gin.Context) {
	sroom.Reboot(c)
}

// post отправки ссылки
func LinkRoom(c *gin.Context) {
	sroom.Link(c)
}

// post блокировки комнаты скприн
func LSRoom(c *gin.Context) {
	sroom.Ls(c)
}

func DelRoom(c *gin.Context) {
	sroom.DelRoom(c)
}
