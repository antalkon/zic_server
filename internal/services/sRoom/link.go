package sroom

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/antalkon/zic_server/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/levigross/grequests"
)

func Link(c *gin.Context) {
	id := c.Param("id")

	// Привязываем JSON данные к модели
	var s models.OpenLink
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Получаем IP-адреса
	lip, err := roompg.GetOnPCsLIPs(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "no IPs found or error retrieving IPs"})
		return
	}

	// Настройка запроса
	options := &grequests.RequestOptions{
		Headers: map[string]string{
			"User-Agent":    "PostmanRuntime/7.42.0",
			"Postman-Token": uuid.New().String(),
			"Connection":    "keep-alive",
		},
		JSON:           s, // Передаем JSON данные
		RequestTimeout: 6 * time.Second,
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	success, bad := 0, 0

	for _, ip := range lip {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			resp, err := grequests.Post(fmt.Sprintf("http://%s:3000/pc/link", ip), options)
			mu.Lock()
			defer mu.Unlock()
			if err != nil || resp.StatusCode != http.StatusOK {
				bad++
			} else {
				success++
			}
			if resp != nil {
				resp.Close()
			}
		}(ip)
	}
	wg.Wait()

	// Возвращаем результаты
	c.JSON(200, gin.H{"success": success, "bad": bad, "message": "Ссылка отправлена на все доступные ПК!"})
}
