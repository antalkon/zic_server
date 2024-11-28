package ssl

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RunWithTLS запускает сервер с поддержкой HTTPS
func RunWithTLS(r *gin.Engine, address, certPath, keyPath string) error {

	// Проверяем, заданы ли пути к сертификату и ключу
	if certPath == "" || keyPath == "" {
		return fmt.Errorf("пути к SSL-сертификату и ключу не заданы")
	}

	// Запуск сервера с использованием TLS
	fmt.Printf("Сервер запущен с SSL на %s\n", address)
	return http.ListenAndServeTLS(address, certPath, keyPath, r)
}
