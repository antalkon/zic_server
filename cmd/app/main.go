package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/antalkon/zic_server/internal/router"
	"github.com/antalkon/zic_server/pkg/config"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/antalkon/zic_server/pkg/tools"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация логгера
	logger.InitLogger()

	// Загрузка конфигурации
	if err := config.LoadServerConfig(); err != nil {
		logger.LogError(err)
		fmt.Println("Ошибка загрузки конфигурации. Проверьте логи.")
		return
	}

	// Выводим адрес сервера из конфигурации для отладки
	address := viper.GetString("http_server.address")
	if address == "" {
		logger.LogError(errors.New("Проблема с адресом серера"))
		fmt.Println("Адрес сервера не задан в конфигурациыи. Проверьте настройки.")
		return
	}
	fmt.Printf("Запуск сервера на %s\n", address)

	// Настраиваем Gin с CORS middleware
	r := gin.Default()

	// Добавляем CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Обрабатываем OPTIONS-запросы (CORS preflight)
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Настраиваем маршруты
	router.SetupRoutes(r)

	// Проверяем, нужно ли активировать базу данных
	if viper.GetBool("activate") {
		if err := pggorm.InitDB(); err != nil {
			logger.LogFatal("Ошибка инициализации базы данных: %v", err)
			return
		}
		fmt.Println("База данных успешно инициализирована")
		go tools.LoadCPU()
		go tools.LoadRAM()
		go tools.LoadNetwork()
		go tools.CheckStatus()
	} else {
		fmt.Println("Активация базы данных отключена в конфигурации")
	}

	// Запускаем сервер на указанном адресе
	if err := r.Run(address); err != nil {
		logger.LogError(fmt.Errorf("Ошибка запуска сервера: %w", err))
		fmt.Printf("Ошибка при запуске сервера на %s: %v\n", address, err)
	}
}
