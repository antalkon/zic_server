package handler

import (
	sroom "github.com/antalkon/zic_server/internal/services/sRoom"
	"github.com/gin-gonic/gin"
)

func AddNewRoom(c *gin.Context) {
	sroom.NewRoom(c)
}
