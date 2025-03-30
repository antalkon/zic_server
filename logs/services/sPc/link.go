package spc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func PcLink(c *gin.Context) {
	lip := c.Param("lip")

	// Привязываем JSON данные к модели
	var l models.OpenLink
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Создаем JSON тело запроса
	jsonData, err := json.Marshal(map[string]string{"url": l.URL})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode JSON"})
		return
	}

	// Отправляем POST запрос с JSON телом
	resp, err := http.Post(fmt.Sprintf("http://%s:3000/pc/link", lip), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send link request to PC: %v", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "PC link request failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link request sent successfully"})
}
