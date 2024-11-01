package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Сохраняем загрузку ОЗУ в таблицу
func NewRAMRecord(usage float64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is nil")
	}
	ramLoad := models.RAMLoad{
		Timestamp: time.Now(),
		Usage:     usage,
	}
	return db.Create(&ramLoad).Error
}
