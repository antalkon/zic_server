package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Обнавляем данный ПК в таблице по ID
func UpdatePcById(id uint64, updatedPc *models.Computer) error {
	db := pggorm.DB
	if db == nil {
		err := errors.New("database connection is not established")
		return err
	}
	if err := db.Model(&models.Computer{}).
		Where("id = ?", id).
		Updates(updatedPc).Error; err != nil {
		return err
	}
	return nil
}
