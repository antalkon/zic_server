package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получаем данный по lip
func GetPcData(lip string) (models.Computer, error) {
	db := pggorm.DB
	if db == nil {
		return models.Computer{}, errors.New("DB connection is nil")
	}
	var pc models.Computer
	if err := db.Model(&models.Computer{}).
		Where("l_ip = ?", lip).
		First(&pc).Error; err != nil {
		return pc, err
	}
	return pc, nil
}

// Получаем данный ПК по ID из таблицы
func GetPcDataId(id string) (models.Computer, error) {
	db := pggorm.DB
	if db == nil {
		return models.Computer{}, errors.New("DB connection is nil")
	}
	var pc models.Computer
	if err := db.Model(&models.Computer{}).
		Where("id = ?", id).
		First(&pc).Error; err != nil {
		return pc, err
	}
	return pc, nil
}
