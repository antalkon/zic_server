package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получаем загрузку сети за 24 часа из таблицы
func NetworkLoad24() ([]models.NetworkLoad, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("error loading database")
	}
	var networkData []models.NetworkLoad
	cutoffTime := time.Now().Add(-24 * time.Hour)
	result := db.Where("timestamp >= ?", cutoffTime).Order("timestamp ASC").Find(&networkData)
	if result.Error != nil {
		return nil, result.Error
	}
	return networkData, nil
}
