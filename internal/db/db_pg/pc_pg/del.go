package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Удаялем ПК из таблицы по ID
func DeletePcById(id string) error {
	db := pggorm.DB
	if db == nil {
		err := errors.New("database connection is not established")
		return err
	}
	db.Delete(&models.Computer{}, id)
	return nil
}
