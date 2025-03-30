package ssettings

import (
	"net/http"

	"github.com/antalkon/zic_server/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Version(c *gin.Context) {
	err := config.LoadServerConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	version := viper.GetString("version")

	c.JSON(200, gin.H{
		"success": true,
		"version": version,
	})

}
