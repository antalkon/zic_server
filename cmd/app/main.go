package main

import (
	"fmt"

	"github.com/antalkon/zic_server/internal/router"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load inital function
	logger.InitLogger()

	// Load configuration function

	if err := config.LoadServerConfig(); err != nil {
		logger.LogError(err)
		fmt.Println("Ошибка загрузки конфигурации. Проверьте логи.")
		return
	}
	activate := viper.GetBool("activate")
	_ = activate
	r := gin.Default()

	// Настраиваем маршруты
	router.SetupRoutes(r)

	// Запускаем go
	address := viper.GetString("http_server.adress")
	if err := r.Run(address); err != nil {
		logger.LogError(err)
		fmt.Println("server started:", address)
	}

	// Check activation status server
	if viper.GetBool("activate") {

		logger.LogInfo("Сервер активирован")
		fmt.Println("Server activated")

		return
	}

}
