package spc

import (
	"fmt"
	"net/http"

	pcpg "github.com/antalkon/zic_server/internal/db/db_pg/pc_pg"
	"github.com/gin-gonic/gin"
)

func PcBlock(c *gin.Context) {
	id := c.Param("id")

	lip, err := pcpg.ChangeStatus(id, "blocked")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:3000/pc/block", lip))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Errorf("failed to change status: %w", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(500, gin.H{"error": "Failed to block PC"})
		return
	}

	c.JSON(200, gin.H{"message": "PC blocked and server request sent"})

}

func PcUnBlock(c *gin.Context) {
	id := c.Param("id")

	lip, err := pcpg.ChangeStatus(id, "work")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:3000/pc/unblock", lip))
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Errorf("failed to change status: %w", err)})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(500, gin.H{"error": "Failed to block PC"})
		return
	}

	c.JSON(200, gin.H{"message": "PC unblocked and server request sent"})

}
