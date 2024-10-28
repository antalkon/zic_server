package router

import (
	"fmt"

	"github.com/antalkon/zic_server/internal/handler"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/antalkon/zic_server/pkg/sessions"
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
	r.Use(sessions.InitSessionStore())

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/public/**/*")

	main := r.Group("/")
	{
		if viper.GetBool("activate") {
			main.GET("/", handler.LoginPage)
			main.GET("/dashboard", handler.Dashboard)
			main.GET("/dashboard/settings", handler.SettingsPage)
			main.GET("/dashboard/load/cpu", handler.LoadCPUPage)
			main.GET("/dashboard/load/ram", handler.LoadRAMPage)
			main.GET("/dashboard/load/network", handler.LoadNetworkPage)
			main.GET("/dashboard/load/storage", handler.LoadStoragePage)

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
		settingsApi.POST("/sys/status", handler.SysStaus)
		settingsApi.POST("/tg", handler.TgData)
		settingsApi.PUT("/tg", handler.TgNewData)
		settingsApi.POST("/sec/network", handler.SecNetwork)
		settingsApi.POST("/sec/firewall", handler.SecFirewall)
		settingsApi.POST("/sec/auth", handler.SecAuth)
		settingsApi.POST("/sec/settings", handler.SecGet)

		settingsApi.GET("/info", handler.InfoVesion)

	}

	tools := r.Group("/tools/api")
	{
		tools.POST("/restart", handler.Restart)
		tools.POST("/load/cpu", handler.LoadCPU)
		tools.POST("/load/ram", handler.LoadRAM)
		tools.POST("/load/network", handler.LoadNetwork)
		tools.POST("/load/storage", handler.LoadStorage)

	}
	// Группа маршрутов для компьютеров
	pcApi := r.Group("/pc/api")
	{
		pcApi.POST("/new", handler.AddNewPc)
		pcApi.POST("/count", handler.PcCount)

		// pcApi.POST("/coment/:id")
	}

	roomApi := r.Group("/room/api")
	{
		roomApi.POST("/new", handler.AddNewRoom)
		roomApi.POST("/count", handler.RoomCount)

	}

	authApi := r.Group("/auth/api")
	{
		authApi.POST("/login", handler.Login)
	}
}
