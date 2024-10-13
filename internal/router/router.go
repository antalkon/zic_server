package router

import (
	"fmt"

	"github.com/antalkon/zic_server/internal/handler"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(r *gin.Engine) {
	if err := config.LoadServerConfig(); err != nil {
		logger.LogError(err)
		fmt.Println("Ошибка загрузки конфигурации. Проверьте логи.")
		return
	}

	// Включаем стандартное логирование запросов и обработку ошибок
	r.Use(gin.Logger())   // Логирование всех запросов
	r.Use(gin.Recovery()) // Восстановление после паники и логирование ошибок

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/public/**/*")

	main := r.Group("/")
	{
		if viper.GetBool("activate") {

		} else {
			main.GET("/", handler.ActivatePage)
		}

	}

	// Группа маршрутов для пользователей
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", handler.GetUsers)
		userGroup.POST("/", handler.CreateUser)
	}

	// Группа маршрутов для настроек
	settingsApi := r.Group("/setting/api")
	{
		settingsApi.POST("/activate", handler.Activate)
	}

	// Группа маршрутов для компьютеров
	pcApi := r.Group("/pc/api")
	{
		pcApi.POST("/new", handler.AddNewPc)
		// pcApi.POST("/coment/:id")
	}

	roomApi := r.Group("/room/api")
	{
		roomApi.POST("/new", handler.AddNewRoom)
	}
}
