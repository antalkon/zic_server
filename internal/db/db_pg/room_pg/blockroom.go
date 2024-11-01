package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Изменяем статус ПК в команате
func ChangePCRoomStatus(roomID string, status string) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("DB connection is nil")
	}
	if err := db.Model(&models.Computer{}).Where("room_id = ?", roomID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
