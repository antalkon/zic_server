package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func CheckedTask(id uint64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is not established")
	}
	var task models.DeliveredTasks
	if err := db.Where("id =?", id).First(&task).Error; err != nil {
		return err
	}
	task.Checked = true
	if err := db.Save(&task).Error; err != nil {
		return err
	}
	return nil
}
