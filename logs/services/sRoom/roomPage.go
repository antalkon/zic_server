package sroom

import (
	roompg "github.com/antalkon/zic_server/internal/db/db_pg/room_pg"
	"github.com/gin-gonic/gin"
)

func RoomPage(c *gin.Context) {
	id := c.Param("id")

	// Получение количества ПК с room_id = id и on = true
	on, err := roompg.GetCountPCOn(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получение количества ПК с room_id = id и on = false
	off, err := roompg.GetCountPCOff(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получение общего количества ПК с room_id = id
	count, err := roompg.GetTotalPCCount(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получение количества ПК с room_id = id и status = "blocked"
	blocked, err := roompg.GetCountPCBlocked(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получение количества ПК с room_id = id и status = "work"
	worked, err := roompg.GetCountPCWorking(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Получение массива всех ПК с room_id = id
	pc, err := roompg.GetAllPCsInRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Рендеринг HTML шаблона с данными
	c.HTML(200, "room.html", gin.H{
		"title":   "Страница комнаты",
		"on":      on,
		"off":     off,
		"count":   count,
		"blocked": blocked,
		"worked":  worked,
		"allPc":   pc,
		"id":      id,
	})
}
