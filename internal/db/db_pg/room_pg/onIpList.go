package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetOnPCsLIPs(room string) ([]string, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("DB connection is nil")
	}

	var ips []string
	if err := db.Model(&models.Computer{}).
		Where("room_id = ? AND \"on\" = ?", room, true).
		Pluck("l_ip", &ips).Error; err != nil {
		return nil, err
	}

	return ips, nil
}
