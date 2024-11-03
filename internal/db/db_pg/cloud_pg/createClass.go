package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func CreateClass(c *models.Class) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database CreateClass")
	}

	if err := db.Create(c).Error; err != nil {
		return err
	}
	return nil
}
