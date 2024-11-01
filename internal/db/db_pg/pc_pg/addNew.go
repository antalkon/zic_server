package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Добавляем запись нового пк в таблицу
func AddNewPc(pc *models.Computer) error {
	db := pggorm.DB
	if db == nil {
		err := errors.New("database connection is not established")
		return err
	}
	if err := db.Create(pc).Error; err != nil {
		return err
	}
	return nil
}
