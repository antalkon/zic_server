package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получаем LIP всех пк в команате
func GetAllPcIpInRoom(id string) ([]string, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("error loading database")
	}
	var ips []string
	if err := db.Model(&models.Computer{}).
		Where("room_id = ?", id).
		Pluck("l_ip", &ips).Error; err != nil {
		return nil, err
	}
	return ips, nil
}
