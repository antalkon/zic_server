package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetAllClass() ([]models.Class, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("Error creating database GetAllClass")
	}
	var c []models.Class
	if err := db.Find(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}
