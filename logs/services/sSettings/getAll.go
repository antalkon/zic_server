package ssettings

import (
	"net/http"

	"github.com/antalkon/zic_server/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Get(c *gin.Context) {
	err := config.LoadDataConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load configuration"})
		return
	}

	firewall := viper.GetBool("sec.firewall")
	gbnet := viper.GetBool("sec.gbnet")
	reqauth := viper.GetBool("sec.reqauth")

	c.JSON(200, gin.H{
		"success": true,
		"fw":      firewall,
		"gbn":     gbnet,
		"ra":      reqauth,
		"info":    "Успешно!",
	})
}
