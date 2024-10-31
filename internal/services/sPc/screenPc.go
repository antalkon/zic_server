package spc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Screen(c *gin.Context) {
	lip := c.Param("lip")

	// Отправляем запрос на http://{lip}:3000/pc/screen
	url := fmt.Sprintf("http://%s:3000/pc/screen", lip)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch screen from remote PC"})
		return
	}
	defer resp.Body.Close()

	// Парсим имя из заголовка JSON
	var jsonResp struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	// Проверяем и создаем папку, если нужно
	outputDir := "data/screens"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create output directory"})
		return
	}

	// Формируем путь для сохранения файла
	filename := fmt.Sprintf("%s.png", jsonResp.Name)
	filepath := filepath.Join(outputDir, filename)

	// Открываем файл для записи
	file, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create output file"})
		return
	}
	defer file.Close()

	// Копируем содержимое тела запроса (изображение) в файл
	if _, err := io.Copy(file, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save image"})
		return
	}

	// Возвращаем ссылку на сохраненное изображение
	c.JSON(http.StatusOK, gin.H{
		"screen": fmt.Sprintf("/%s/%s", outputDir, filename),
	})
}
