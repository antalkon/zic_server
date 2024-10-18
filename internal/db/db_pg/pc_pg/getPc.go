package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
	"github.com/antalkon/zic_server/pkg/logger"
)

func GetPcCountAndOn() (error, int64, int64) {
	db := pggorm.DB
	if db == nil {
		logger.LogError(errors.New("Error creating database GetPcCountAndOn"))
		return errors.New("Error connect to database"), 0, 0
	}

	var totalCount int64
	var onCount int64

	// Получаем общее количество ПК
	if err := db.Model(&models.Computer{}).Count(&totalCount).Error; err != nil {
		logger.LogError(err)
		return err, 0, 0
	}

	// Получаем количество ПК, где on = true
	if err := db.Model(&models.Computer{}).Where("\"on\" = ?", true).Count(&onCount).Error; err != nil {
		logger.LogError(err)
		return err, 0, 0
	}

	return nil, totalCount, onCount
}
