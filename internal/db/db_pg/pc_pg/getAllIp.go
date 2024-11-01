package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

type IP struct {
	LIP string `json:"lip"`
}

// Получаем все LIP адреса всез ПК
func GetAllIp() ([]string, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("DB connection is not established")
	}
	var ips []string
	if err := db.Model(&models.Computer{}).Pluck("l_ip", &ips).Error; err != nil {
		return nil, err
	}
	return ips, nil
}
