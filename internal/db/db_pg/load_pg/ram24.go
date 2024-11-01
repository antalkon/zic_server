package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получаем загрузку RAM из таблицы за 24 часа
func RamLoad24() ([]models.RAMLoad, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("Error loading database")
	}
	var RramData []models.RAMLoad
	cutoffTime := time.Now().Add(-24 * time.Hour)
	result := db.Where("timestamp >= ?", cutoffTime).Order("timestamp ASC").Find(&RramData)
	if result.Error != nil {
		return nil, result.Error
	}
	return RramData, nil
}
