package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получаем количество ПК и количество включенных ПК
func GetPcCountAndOn() (error, int64, int64) {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error connect to database"), 0, 0
	}
	var totalCount int64
	var onCount int64
	if err := db.Model(&models.Computer{}).Count(&totalCount).Error; err != nil {
		return err, 0, 0
	}
	if err := db.Model(&models.Computer{}).Where("\"on\" = ?", true).Count(&onCount).Error; err != nil {
		return err, 0, 0
	}
	return nil, totalCount, onCount
}
