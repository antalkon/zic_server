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
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Инициализация логгера
	conf := logger.Config{
		Filename:   "logs/server.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
		Level:      logrus.InfoLevel,
	}
	log, err := logger.InitLogger(conf)
	if err != nil {
		fmt.Printf("Ошибка инициализации логгера: %v\n", err)
		return
	}

	// Загрузка конфигурации
	if err := config.LoadServerConfig(); err != nil {
		log.Error(errors.New("Ошибка загрузки конфигурации"))
		fmt.Println("Ошибка загрузки конфигурации. Проверьте логи.")
		return
	}

	// Получаем адрес сервера из конфигурации для отладки
	address := viper.GetString("http_server.address")
	if address == "" {
		log.Error(errors.New("Проблема с адресом сервера"))
		return
	}
	fmt.Printf("Запуск сервера на %s\n", address)

	// Настройка роутеров и gin
	config.SetupGinMode(log)
	r := gin.Default()
	router.SetupRoutes(r, log)

	// Если сервер активирован
	if viper.GetBool("activate") {
		// Инициализация базы данных
		if err := pggorm.InitDB(); err != nil {
			log.Fatal(err)
			return
		}

		// Запуск служб
		go tools.LoadCPU()
		go tools.LoadRAM()
		go tools.LoadNetwork()
		go tools.CheckStatus()
	} else {
		fmt.Println("!!! --- СЕРВЕР НЕ АКТИВИРОВАН --- !!!")
	}

	// Запуск сервера на адресе
	if err := r.Run(address); err != nil {
		log.Error(fmt.Errorf("Ошибка запуска сервера: %w", err))
	}
}
