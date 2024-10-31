package spc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/antalkon/zic_server/internal/models"
	getip "github.com/antalkon/zic_server/pkg/getIP"
	"github.com/antalkon/zic_server/pkg/logger"
	"github.com/antalkon/zic_server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddPc(c *gin.Context) {
	var pc models.Computer

	// Привязка JSON данных к модели
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Генерация случайного ID
	id, err := utils.GenerateRandomNumber(12)
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to generate random ID"})
		fmt.Println(err)
		return
	}
	pc.ID = uint64(id)

	// Добавление ПК в базу данных
	if err := pcpg.AddNewPc(&pc); err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)

		return
	}

	// Получаем локальный и публичный IP
	pip, lip, err := getip.GetIps()
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to get IPs"})
		fmt.Println(err)

		return
	}

	// Подготовка данных для отправки на http://{pc.LIP}:3000/server/invite
	inviteData := map[string]interface{}{
		"local_ip":  lip,
		"public_ip": pip,
		"port":      8385,
	}
	jsonData, err := json.Marshal(inviteData)
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to encode JSON"})
		fmt.Println(err)

		return
	}

	// Формирование URL и отправка POST запроса
	url := fmt.Sprintf("http://%s:3000/server/invite", pc.LIP)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to create request"})
		fmt.Println(err)

		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.LogError(err)
		c.JSON(500, gin.H{"error": "Failed to send request"})
		fmt.Println(err)

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(500, gin.H{"error": "Failed to invite PC"})
		fmt.Println(err)

		return
	}

	c.JSON(200, gin.H{"success": "PC added and invited successfully"})
}
