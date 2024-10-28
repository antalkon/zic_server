package main

import (
	"errors"
	"fmt"

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
	address := viper.GetString("http_server.address") // исправил "adress" на "address"
	if address == "" {
		logger.LogError(errors.New("Проблема с адресом серера"))
		fmt.Println("Адрес сервера не задан в конфигурациыи. Проверьте настройки.")
		return
	}
	fmt.Printf("Запуск сервера на %s\n", address)

	// Настраиваем маршруты с использованием gin
	r := gin.Default()
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
	} else {
		fmt.Println("Активация базы данных отключена в конфигурации")
	}

	// Запускаем сервер на указанном адресе
	if err := r.Run(address); err != nil {
		logger.LogError(fmt.Errorf("Ошибка запуска сервера: %w", err))
		fmt.Printf("Ошибка при запуске сервера на %s: %v\n", address, err)
	}
}
