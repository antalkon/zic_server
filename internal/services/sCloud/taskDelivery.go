package scloud

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	cloudpg "github.com/antalkon/zic_server/internal/db/db_pg/cloud_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func TaskDelivery(c *gin.Context) {
	var td models.DeliveredTasks
	data := c.PostForm("data")
	if err := json.Unmarshal([]byte(data), &td); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		fmt.Println("Ошибка обработки JSON:", err)
		return
	}

	// Получаем файл из запроса
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Ошибка получения файла из запроса:", err)
		c.JSON(400, gin.H{"error": "Не удалось загрузить файл"})
		return
	}

	// Генерируем уникальный ID и формируем имя файла
	uniqueID := GenerateRandomIDInt64() // Генерируем уникальный ID
	fileName := fmt.Sprintf("%d%s", uniqueID, filepath.Ext(file.Filename))

	// Устанавливаем ID и имя файла для DeliveredTasks
	td.ID = uint64(uniqueID) // ID без расширения
	td.File = fileName       // Полное имя файла с расширением

	// Создаем директорию для хранения файла
	dirPath := filepath.Join("data", "dtask")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("Ошибка создания директории:", err)
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка создания директории %s", dirPath)})
		return
	}

	// Сохраняем файл
	savePath := filepath.Join(dirPath, fileName)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка сохранения файла в %s", savePath)})
		return
	}

	// Сохраняем задание в базе данных
	if err := cloudpg.SaveDeliveryTasl(&td); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Файл успешно загружен и данные сохранены"})
}
