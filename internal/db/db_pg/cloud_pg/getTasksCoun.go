package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetTasksCount() (int64, int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, 0, errors.New("database connection is not established")
	}
	var tasks int64
	var pending int64
	if err := db.Model(&models.DeliveredTasks{}).Count(&tasks).Error; err != nil {
		return 0, 0, err
	}
	if err := db.Model(&models.DeliveredTasks{}).Where("checked = false").Count(&pending).Error; err != nil {
		return 0, 0, err
	}
	return tasks, pending, nil
}
