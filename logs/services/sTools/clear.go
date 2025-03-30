package stools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	toolspg "github.com/antalkon/zic_server/internal/db/db_pg/tools_pg"
	"github.com/gin-gonic/gin"
)

func Clear(c *gin.Context) {
	if err := toolspg.ClearCpuTable(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := toolspg.ClearNetworkTable(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := toolspg.ClearRamTable(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := toolspg.ClearDTasksTable(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := toolspg.ClearPTasksTable(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	err := clearDirectories()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "Clearing successful"})
}

func clearDirectories() error {
	directories := []string{"./data/class", "./data/dtask"}

	for _, dir := range directories {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("не удалось прочитать директорию %s: %w", dir, err)
		}

		for _, file := range files {
			itemPath := filepath.Join(dir, file.Name())
			if err := os.RemoveAll(itemPath); err != nil {
				return fmt.Errorf("не удалось удалить %s: %w", itemPath, err)
			}
		}
		fmt.Printf("Директория %s успешно очищена\n", dir)
	}

	return nil
}
