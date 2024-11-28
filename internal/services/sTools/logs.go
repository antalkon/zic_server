package stools

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetLogs(c *gin.Context) {
	// Получаем количество строк из параметров запроса
	linesParam := c.DefaultQuery("lines", "300") // По умолчанию 100 строк
	lines, err := strconv.Atoi(linesParam)
	if err != nil || lines <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'lines' parameter"})
		return
	}

	// Путь к файлу логов
	logFilePath := "logs/server.log"

	// Читаем последние строки из логов
	logs, err := readLogsChunk(logFilePath, lines)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем логи в JSON-формате
	c.JSON(http.StatusOK, gin.H{"logs": logs})
}

// readLogsChunk читает последние N строк из файла
func readLogsChunk(filePath string, lines int) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Создаем слайс для хранения последних строк
	var result []string
	buffer := make([]byte, 4096) // Читаем файл порциями по 4 КБ

	// Позиционируемся в конец файла
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()

	offset := size
	for offset > 0 && len(result) < lines {
		if offset < int64(len(buffer)) {
			buffer = make([]byte, offset)
		}

		// Смещаемся назад и читаем блок
		offset -= int64(len(buffer))
		if _, err := file.ReadAt(buffer, offset); err != nil && err != io.EOF {
			return nil, err
		}

		// Разбиваем блок на строки
		linesInChunk := strings.Split(string(buffer), "\n")
		for i := len(linesInChunk) - 1; i >= 0 && len(result) < lines; i-- {
			if linesInChunk[i] != "" {
				result = append(result, linesInChunk[i])
			}
		}
	}

	// Переворачиваем результат, чтобы строки были в хронологическом порядке
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, nil
}
