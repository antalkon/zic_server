package scloud

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	var m models.Class
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Генерация уникального FolderID для класса
	m.FolderID = GenerateRandomString()

	// Сохранение класса в базе данных
	if err := cloudpg.CreateClass(&m); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Создание директории для класса с именем FolderID
	dirPath := filepath.Join("data", "class", m.FolderID)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка создания директории для класса: %s", err)})
		return
	}

	c.JSON(200, gin.H{"success": true})
}

// GenerateRandomString генерирует случайную строку длиной 8 символов
func GenerateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // Инициализация случайного генератора
	result := make([]byte, 8)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
