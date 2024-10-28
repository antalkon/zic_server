package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func CpuLoad24() ([]models.CPULoad, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("Error loading database")
	}

	var cpuData []models.CPULoad
	cutoffTime := time.Now().Add(-24 * time.Hour)

	// Выполняем запрос к базе данных для получения данных за последние 24 часа
	result := db.Where("timestamp >= ?", cutoffTime).Order("timestamp ASC").Find(&cpuData)
	if result.Error != nil {
		return nil, result.Error
	}

	// Возвращаем данные
	return cpuData, nil
}
