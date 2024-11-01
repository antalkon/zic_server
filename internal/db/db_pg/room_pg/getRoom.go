package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetRoom() (error, int64, int64) {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connect to database"), 0, 0
	}

	var totalCount int64
	var onCount int64

	// Получаем общее количество ПК
	if err := db.Model(&models.Room{}).Count(&totalCount).Error; err != nil {
		return err, 0, 0
	}

	// Получаем количество ПК, где on = true
	if err := db.Model(&models.Room{}).Where("\"on\" = ?", true).Count(&onCount).Error; err != nil {
		return err, 0, 0
	}

	return nil, totalCount, onCount
}
