package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetRooms() ([]models.Room, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("error loading database")
	}

	var rooms []models.Room

	// Выполняем запрос для получения списка комнат
	if err := db.Find(&rooms).Error; err != nil {
		return nil, err
	}

	return rooms, nil
}
