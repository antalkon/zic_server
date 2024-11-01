package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
	"github.com/antalkon/zic_server/pkg/logger"
)

func UpdatePcById(id uint64, updatedPc *models.Computer) error {
	db := pggorm.DB
	if db == nil {
		err := errors.New("database connection is not established")
		logger.LogError(err)
		return err
	}

	// Поиск ПК по ID и обновление всех полей
	if err := db.Model(&models.Computer{}).
		Where("id = ?", id).
		Updates(updatedPc).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
