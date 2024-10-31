package spc

import (
	"fmt"
	"time"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/levigross/grequests"
)

func RebootPc(c *gin.Context) {
	id := c.Param("id")

	ip, err := pcpg.GetIpByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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

	url := fmt.Sprintf("http://%s:3000/pc/reboot", ip)
	resp, err := grequests.Get(url, options)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(500, gin.H{"error": "failed to turn off PC"})
		return
	}

	c.JSON(200, gin.H{"message": "PC turned off"})

}
