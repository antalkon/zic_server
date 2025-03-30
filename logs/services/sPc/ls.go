package spc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PcLS(c *gin.Context) {
	lip := c.Param("lip")

	// Привязываем JSON данные к модели

	// Отправляем POST запрос с JSON телом
	resp, err := http.Get(fmt.Sprintf("http://%s:3000/pc/ls", lip))
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
