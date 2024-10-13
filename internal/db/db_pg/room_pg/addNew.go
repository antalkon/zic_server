package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func NewRoom(r *models.Room) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database NewRoom")
	}

	if err := db.Create(r).Error; err != nil {
		return err
	}
	return nil
}
