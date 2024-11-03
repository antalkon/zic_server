package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func DelTaskDel(id uint64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is not established")
	}
	db.Delete(&models.DeliveredTasks{}, id)
	return nil
}
