package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func DelClass(id string) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is not established")
	}
	db.Delete(&models.Class{}, id)
	return nil
}
