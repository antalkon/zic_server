package loadpg

import (
	"errors"
	"time"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Сохраняем загрузки сети в таблицу
func NewNetworkRecord(bytesRecv, bytesSent uint64) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is nil")
	}
	networkLoad := models.NetworkLoad{
		Timestamp: time.Now(),
		BytesRecv: bytesRecv,
		BytesSent: bytesSent,
	}
	return db.Create(&networkLoad).Error
}
