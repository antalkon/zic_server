package stools

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLogsStats(c *gin.Context) {
	// Путь к файлу логов
	logFilePath := "logs/server.log"

	// Получаем статистику
	stats, err := getLogStats(logFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Отправляем результат клиенту
	c.JSON(http.StatusOK, stats)
}

func getLogStats(filePath string) (map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Инициализируем статистику
	stats := map[string]int{
		"info":  0,
		"warn":  0,
		"error": 0,
		"fatal": 0,
	}

	// Сканируем файл построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "level=info") {
			stats["info"]++
		} else if strings.Contains(line, "level=warn") {
			stats["warn"]++
		} else if strings.Contains(line, "level=error") {
			stats["error"]++
		} else if strings.Contains(line, "level=fatal") {
			stats["fatal"]++
		}
	}

	// Проверяем ошибки сканера
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
