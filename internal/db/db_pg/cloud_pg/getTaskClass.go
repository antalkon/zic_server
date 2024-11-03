package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetClassTasks(id string) ([]models.PublicTasks, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("database connection is not established")
	}
	var tasks []models.PublicTasks
	if err := db.Table("public_tasks").Where("classes =?", id).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
