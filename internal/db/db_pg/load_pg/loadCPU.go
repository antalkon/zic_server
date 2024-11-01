package loadpg

import (
	"errors"
	"log"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Создание запись нагрузки CPU
func NewCPURecord(usage []float64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database connection")
	}
	cpuLoad := models.CPULoad{
		Timestamp: time.Now(),
		Usage:     usage[0],
	}
	if err := db.Create(&cpuLoad).Error; err != nil {
		log.Printf("Error saving CPU usage to database: %v", err)
	}
	return nil

}
