package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func SaveNewTask(t *models.PublicTasks) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database SaveNewTask")
	}
	if err := db.Create(&t).Error; err != nil {
		return err
	}
	return nil

}
