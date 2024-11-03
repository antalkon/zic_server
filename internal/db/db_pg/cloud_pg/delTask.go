package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func DelTask(task uint64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database DelTask")
	}
	if err := db.Delete(&models.PublicTasks{}, task).Error; err != nil {
		return err
	}
	return nil
}
