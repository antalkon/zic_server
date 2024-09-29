package router

import (
	"github.com/antalkon/zic_server/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handler.GetUsers)
		userGroup.POST("/", handler.CreateUser)
	}

	settingsApi := r.Group("/setting/api")
	{
		settingsApi.POST("/activate", handler.Activate)

	}
}
