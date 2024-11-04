package sactivate

import (
	"fmt"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	getip "github.com/antalkon/zic_server/pkg/getIP"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Activation handles the activation process.
func Activation(c *gin.Context) {
	var sa models.SActivate
	var sar models.SActivateRequest
	// Parsing JSON from request
	if err := c.ShouldBindJSON(&sa); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}
	// Getting public and local IPs
	pIP, lIP, err := getip.GetIps()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get IP addresses: " + err.Error()})
		return
	}
	// Loading and setting DB config
	if err := loadAndSetConfig(sa); err != nil {
		c.JSON(500, gin.H{"error": "Failed to configure the server"})
		return
	}
	// Prepare request data
	sar.PIP = pIP
	sar.LIP = lIP
	sar.Version = "1.0 beta"
	sar.User = sa.User
	sar.Password = sa.Password
	sar.Port = sa.Port
	// Send activation request
	if err := SendActivationRequest(sa.Code, sar); err != nil {
		c.JSON(500, gin.H{"error": "Failed to send activation request: " + err.Error()})
		return
	}
	err = SetConfig()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to set configuration"})
		return
	}
	c.JSON(200, gin.H{"status": "Activation successful"})
}

// loadAndSetConfig loads the config and applies changes.
func loadAndSetConfig(sa models.SActivate) error {
	if err := config.LoadDBConfig(); err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}
	// Set new configuration
	viper.Set("port", sa.Port)
	viper.Set("user", sa.User)
	viper.Set("password", sa.Password)
	// Save configuration back to the file
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config: %v", err)
	}
	return nil
}

func SetConfig() error {
	if err := config.LoadServerConfig(); err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}
	// Set new configuration
	viper.Set("activate", true)
	// Save configuration back to the file
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config: %v", err)
	}
	return nil
}

// SendActivationRequest sends an activation request to the server.
func SendActivationRequest(code string, requestData models.SActivateRequest) error {
	// Serialize request data to JSON
	// jsonData, err := json.Marshal(requestData)
	// if err != nil {
	// 	return fmt.Errorf("error serializing request data: %v", err)
	// }
	// // Define the HTTP client with a timeout
	// client := &http.Client{Timeout: 10 * time.Second}
	// // Send POST request to the server
	// url := fmt.Sprintf("http://zic.zentas.ru/actions/api/v1/add/server/join/%s", code)
	// resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	// if err != nil {
	// 	return fmt.Errorf("error sending POST request: %v", err)
	// }
	// defer resp.Body.Close()
	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response body: %v", err)
	// }
	// // Check if the status code is not OK
	// if resp.StatusCode != http.StatusOK {
	// 	return fmt.Errorf("server returned an error: %s, status code: %d", string(body), resp.StatusCode)
	// }

	return nil
}
