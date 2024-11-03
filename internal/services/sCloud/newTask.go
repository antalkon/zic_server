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

// GenerateRandomIDInt64 генерирует случайный int64 ID
func GenerateRandomIDInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	return 100000 + rand.Int63n(900000) // Генерирует шестизначное случайное число
}

func NewTask(c *gin.Context) {
	// Получение данных из запроса
	title := c.PostForm("title")
	classID := c.PostForm("classID")

	// Проверка, что title и classID получены
	if title == "" || classID == "" {
		c.JSON(400, gin.H{"error": "Отсутствует заголовок или ID класса"})
		return
	}

	// Получение файла из запроса
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Ошибка получения файла из запроса:", err)
		c.JSON(400, gin.H{"error": "Не удалось загрузить файл"})
		return
	}
	fmt.Printf("Получен файл: %s\n", file.Filename)

	// Генерация уникального ID для задания с использованием int64
	taskID := GenerateRandomIDInt64()
	fileName := fmt.Sprintf("%d%s", taskID, filepath.Ext(file.Filename))
	fmt.Printf("Сгенерирован Task ID: %d, имя файла: %s\n", taskID, fileName)

	// Получение FolderID для указанного classID
	folderID, err := cloudpg.GetGolderIDByClassID(classID)
	if err != nil {
		fmt.Println("Ошибка получения FolderID:", err)
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка получения FolderID для класса %s", classID)})
		return
	}
	fmt.Printf("Получен FolderID для класса %s: %s\n", classID, folderID)

	// Создание директории, если она еще не существует
	dirPath := filepath.Join("data", "class", folderID)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println("Ошибка создания директории:", err)
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка создания директории %s", dirPath)})
		return
	}

	// Сохранение файла в папку класса
	savePath := filepath.Join(dirPath, fileName)
	fmt.Printf("Сохранение файла в: %s\n", savePath)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		c.JSON(500, gin.H{"error": fmt.Sprintf("Ошибка сохранения файла в %s", savePath)})
		return
	}

	// Заполнение модели PublicTasks
	task := models.PublicTasks{
		ID:        uint64(taskID), // Преобразуем taskID в uint64
		Classes:   classID,        // Сохраняем как строку
		Title:     title,
		FolderID:  folderID,
		File:      fileName,
		CreatedAt: time.Now(),
	}
	fmt.Println("Заполнение модели PublicTasks:", task)

	// Сохранение задания в базу данных
	if err := cloudpg.SaveNewTask(&task); err != nil {
		fmt.Println("Ошибка сохранения задания в базу данных:", err)
		c.JSON(500, gin.H{"error": "Ошибка сохранения задания в базу данных"})
		return
	}
	fmt.Printf("Задание успешно сохранено в базу данных: Task ID %d\n", taskID)

	c.JSON(200, gin.H{"success": true, "task_id": taskID})
}
