package sroom

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
)

func RoomBlock(c *gin.Context) {
	id := c.Param("id")

	// Заблокировать все ПК с room_id = id в базе данных
	err := roompg.ChangePCRoomStatus(id, "blocked")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получаем все IP-адреса ПК в комнате
	ips, err := roompg.GetAllPcIpInRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	successCount, badCount := 0, 0
	mu := sync.Mutex{} // Мьютекс для безопасного обновления счетчиков

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s:3000/pc/block", ip), nil)
			if err != nil {
				mu.Lock()
				badCount++
				mu.Unlock()
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				mu.Lock()
				badCount++
				mu.Unlock()
				return
			}
			resp.Body.Close()

			mu.Lock()
			successCount++
			mu.Unlock()
		}(ip)
	}

	wg.Wait()

	c.JSON(200, gin.H{
		"message": "ПК в комнате заблокированы",
		"success": successCount,
		"bad":     badCount,
	})
}

func RoomUnBlock(c *gin.Context) {
	id := c.Param("id")

	// Снять блокировку со всех ПК с room_id = id в базе данных
	err := roompg.ChangePCRoomStatus(id, "work")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получаем все IP-адреса ПК в комнате
	ips, err := roompg.GetAllPcIpInRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	successCount, badCount := 0, 0
	mu := sync.Mutex{} // Мьютекс для безопасного обновления счетчиков

	for _, ip := range ips {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
			defer cancel()

			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://%s:3000/pc/unblock", ip), nil)
			if err != nil {
				mu.Lock()
				badCount++
				mu.Unlock()
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				mu.Lock()
				badCount++
				mu.Unlock()
				return
			}
			resp.Body.Close()

			mu.Lock()
			successCount++
			mu.Unlock()
		}(ip)
	}

	wg.Wait()

	c.JSON(200, gin.H{
		"message": "ПК в комнате разблокированы",
		"success": successCount,
		"bad":     badCount,
	})
}
