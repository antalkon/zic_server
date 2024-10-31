package pcpg

import (
	"errors"
	"fmt"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

type IP struct {
	LIP string `json:"lip"`
}

// Функция для получения всех LIP-адресов
func GetAllIp() ([]string, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("DB connection is not established")
	}

	var ips []string

	// Выбираем только поле LIP (l_ip) из таблицы Computer и сохраняем в срезе строк
	if err := db.Model(&models.Computer{}).Pluck("l_ip", &ips).Error; err != nil {
		return nil, err
	}

	return ips, nil
}

func UpdateComputerStatus(ip string, status bool) error {
	db := pggorm.DB
	if db == nil {
		return fmt.Errorf("DB connection is not established")
	}

	// Обновляем статус поля "On" для записи с указанным IP
	if err := db.Model(&models.Computer{}).Where("l_ip = ?", ip).Update("on", status).Error; err != nil {
		return fmt.Errorf("failed to update status for IP %s: %w", ip, err)
	}
	return nil
}
