package spc

import (
	"fmt"
	"net/http"
	"strconv"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
)

func StartPc(c *gin.Context) {
	// Получаем ID из параметра
	id := c.Param("id")

	// Преобразуем ID в int64
	value, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	// Привязываем JSON тело запроса к модели ConnectPc
	var m models.ConnectPc
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON format"})
		return
	}

	// Обновляем IP по ID
	if err := pcpg.UpdateLipById(value, m.Ip); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update PC IP"})
		fmt.Println("Ошибка обновления IP:", err)
		return
	}

	pc, err := pcpg.GetPcDataId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve PC data"})
		fmt.Println("Ошибка получения ПК:", err)
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"status":  pc.Status,
	})
}
