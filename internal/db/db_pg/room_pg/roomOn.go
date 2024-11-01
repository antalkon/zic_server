package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func RoomOn(id string, on bool) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("DB connection is nil")
	}

	// Выполняем обновление поля on для всех ПК с указанным roomID
	if err := db.Model(&models.Room{}).Where("id =?", id).Update("on", on).Error; err != nil {
		return err
	}

	return nil
}
