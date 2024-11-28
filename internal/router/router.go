package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/antalkon/zic_server/internal/handler"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/antalkon/zic_server/pkg/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// RequestLogger middleware для логирования всех запросов
func RequestLogger(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Записываем время начала запроса
		startTime := time.Now()

		// Выполняем следующий middleware в цепочке
		c.Next()

		// Логируем информацию о запросе после выполнения
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		log.Info(fmt.Sprintf("[%d] %s %s | %v", statusCode, method, path, latency))
	}
}

// SetupRoutes настраивает маршруты и принимает логгер для логирования
func SetupRoutes(r *gin.Engine, log logger.Logger) {
	if err := config.LoadServerConfig(); err != nil {
		log.Error(fmt.Errorf("Ошибка загрузки конфигурации: %w", err))
		fmt.Println("Ошибка загрузки конфигурации. Проверьте логи.")
		return
	}

	// Мидлварь для аутентификации сессий
	m := sessions.SessionAuthMiddleware()

	// Используем кастомное логирование запросов
	r.Use(RequestLogger(log))

	// Восстановление после паники и логирование ошибок
	r.Use(gin.Recovery())
	r.Use(sessions.InitSessionStore())

	// Настраиваем статические файлы и шаблоны
	r.Static("/static", "./web/static")
	r.Static("/tasks", "./data/class")
	r.Static("/delivered", "./data/dtask")

	r.LoadHTMLGlob("web/public/**/*")

	// Обработчик маршрутов, если маршрут не найден
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/404")
	})

	// Основная группа
	main := r.Group("/")
	{
		// Маршруты в случае, если сервер активирован
		if viper.GetBool("activate") {
			main.GET("/", handler.LoginPage)
			main.GET("/dashboard", m, handler.Dashboard)
			main.GET("/tasks", handler.TasksPage)
			main.GET("/d", handler.TasksDPage)
			main.GET("/dashboard/settings", m, handler.SettingsPage)
			main.GET("/dashboard/settings/logs", m, handler.SettingsLogsPage)
			main.GET("/dashboard/data", m, handler.DataPage)
			main.GET("/dashboard/cloud", m, handler.CloudPage)
			main.GET("/dashboard/cloud/class", m, handler.CClassPage)
			main.GET("/dashboard/cloud/tasks", m, handler.CTaskPage)
			main.GET("/dashboard/data/room/:id", m, handler.RoomDataPage)
			main.GET("/dashboard/data/pc/:id", m, handler.PcDataPage)
			main.GET("/dashboard/load/cpu", m, handler.LoadCPUPage)
			main.GET("/dashboard/load/ram", m, handler.LoadRAMPage)
			main.GET("/dashboard/load/network", m, handler.LoadNetworkPage)
			main.GET("/dashboard/load/storage", m, handler.LoadStoragePage)
			main.GET("/403", handler.NotAuth)
			main.GET("/404", handler.NotFound)
		} else {
			main.GET("/", handler.ActivatePage)
		}
	}

	// Группа маршрутов для настроек
	settingsApi := r.Group("/setting/api")
	{
		settingsApi.POST("/activate", m, handler.Activate)
		settingsApi.POST("/sys/status", m, handler.SysStaus)
		settingsApi.POST("/tg", m, handler.TgData)
		settingsApi.PUT("/tg", m, handler.TgNewData)
		settingsApi.POST("/sec/network", m, handler.SecNetwork)
		settingsApi.POST("/sec/firewall", m, handler.SecFirewall)
		settingsApi.POST("/sec/auth", m, handler.SecAuth)
		settingsApi.POST("/sec/settings", m, handler.SecGet)
		settingsApi.GET("/info", handler.InfoVesion)
	}

	// Группа для работы с инструментами
	tools := r.Group("/tools/api")
	{
		tools.POST("/load/cpu", m, handler.LoadCPU)
		tools.POST("/load/ram", m, handler.LoadRAM)
		tools.POST("/load/network", m, handler.LoadNetwork)
		tools.POST("/load/storage", m, handler.LoadStorage)
		tools.GET("/clear", m, handler.Clear)
		tools.GET("/logs", m, handler.GetLogs)
		tools.GET("/logs/stats", m, handler.GetLogsStats)

	}

	// Группа маршрутов для компьютеров
	pcApi := r.Group("/pc/api")
	{
		pcApi.POST("/new", m, handler.AddNewPc)
		pcApi.POST("/edit", m, handler.EditPc)
		pcApi.POST("/del/:id", m, handler.DelPC)
		pcApi.POST("/count", m, handler.PcCount)
		pcApi.GET("/ping/server", handler.ServerPing)
		pcApi.POST("/off/:id", m, handler.PcOff)
		pcApi.POST("/reboot/:id", m, handler.PcReboot)
		pcApi.POST("/block/:id", m, handler.PcBlock)
		pcApi.POST("/unblock/:id", m, handler.PcUnBlock)
		pcApi.POST("/start/:id", handler.StartPc)
		pcApi.GET("/screen/:lip", handler.ScreenPc)
		pcApi.POST("/link/:lip", handler.LinkPc)
		pcApi.POST("/ls/:lip", handler.LSPc)
		pcApi.GET("/vnc/:id", handler.VncPC)

	}

	// Группа для комнат
	roomApi := r.Group("/room/api")
	{
		roomApi.POST("/new", m, handler.AddNewRoom)
		roomApi.POST("/count", m, handler.RoomCount)
		roomApi.POST("/room", m, handler.RoomsData)
		roomApi.POST("/room/block/:id", m, handler.BlockRoom)
		roomApi.POST("/room/unblock/:id", m, handler.UnblockRoom)
		roomApi.POST("/room/off/:id", m, handler.OffRoom)
		roomApi.POST("/room/reboot/:id", m, handler.RebootRoom)
		roomApi.POST("/room/link/:id", m, handler.LinkRoom)
		roomApi.POST("/room/ls/:id", m, handler.LSRoom)
	}

	cloudApi := r.Group("/cloud/api")
	{
		cloudApi.POST("/sys/count", m, handler.TasksCount)

		cloudApi.POST("/class", m, handler.CreateClass)
		cloudApi.DELETE("/class/:id", m, handler.DelClass)
		cloudApi.POST("/classes", handler.GetClasses)
		cloudApi.POST("/tasks", m, handler.NewTask)
		cloudApi.DELETE("/tasks/:id", m, handler.DelTask)

		cloudApi.POST("/tasks/all", m, handler.AllTasks)

		cloudApi.POST("/tasks/class/:id", handler.GetTaskClass)

		cloudApi.POST("/tasks/delivery", handler.TaskDelivery)
		cloudApi.GET("/tasks/delivery/d/:id", m, handler.DelTaskDel)
		cloudApi.POST("/delivery/all/:id", m, handler.DeliveryAll)
		cloudApi.GET("/delivery/checked/:id", m, handler.DeliveryChecked)

	}

	// Auth группа
	authApi := r.Group("/auth/api")
	{
		authApi.POST("/login", handler.Login)
	}
}
