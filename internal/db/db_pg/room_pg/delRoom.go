package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func DelRoomById(id string) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is not established")
	}

	// Удаление записи по ID
	if err := db.Where("id = ?", id).Delete(&models.Room{}).Error; err != nil {
		return err
	}

	return nil
}
