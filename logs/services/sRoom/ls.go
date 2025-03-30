package sroom

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/levigross/grequests"
)

func Ls(c *gin.Context) {
	id := c.Param("id")

	lip, err := roompg.GetOnPCsLIPs(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "no IPs found or error retrieving IPs"})
		return
	}

	timeout := 6 * time.Second
	options := &grequests.RequestOptions{
		Headers: map[string]string{
			"User-Agent":      "PostmanRuntime/7.42.0",
			"Accept":          "*/*",
			"Postman-Token":   uuid.New().String(),
			"Host":            "192.168.31.231:3000",
			"Accept-Encoding": "gzip, deflate, br",
			"Connection":      "keep-alive",
		},
		RequestTimeout: timeout,
	}

	// Массивы для успешных и неудачных запросов
	success := 0
	bad := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, ip := range lip {
		wg.Add(1)

		go func(ip string) {
			defer wg.Done()
			url := fmt.Sprintf("http://%s:3000/pc/ls", ip)
			resp, err := grequests.Get(url, options)
			if err != nil || resp.StatusCode != http.StatusOK {
				mu.Lock()
				bad += 1
				mu.Unlock()
				return
			}
			defer resp.Close()

			// Если запрос успешен
			mu.Lock()
			success += 1
			mu.Unlock()
		}(ip)
	}

	// Ожидаем завершения всех запросов
	wg.Wait()

	// Возвращаем результаты
	c.JSON(200, gin.H{
		"success": success,
		"bad":     bad,
	})

}
