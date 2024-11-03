package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetAllTasks() ([]models.PublicTasks, error) {
	var tasks []models.PublicTasks
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("Error creating database GetAllTasks")
	}
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
