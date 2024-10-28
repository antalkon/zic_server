package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func NetworkLoad24() ([]models.NetworkLoad, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("error loading database")
	}

	var networkData []models.NetworkLoad
	cutoffTime := time.Now().Add(-24 * time.Hour)

	// Выполняем запрос к базе данных для получения данных за последние 24 часа
	result := db.Where("timestamp >= ?", cutoffTime).Order("timestamp ASC").Find(&networkData)
	if result.Error != nil {
		return nil, result.Error
	}

	// Возвращаем данные
	return networkData, nil
}
