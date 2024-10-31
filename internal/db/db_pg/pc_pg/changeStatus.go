package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func ChangeStatus(id, status string) (string, error) {
	db := pggorm.DB
	if db == nil {
		return "", errors.New("DB connection is nil")
	}

	// Структура для хранения результата
	var computer models.Computer

	// Обновляем статус и выбираем l_ip
	if err := db.Model(&models.Computer{}).
		Where("id = ?", id).
		Update("status", status).
		Select("l_ip").
		First(&computer).Error; err != nil {
		return "", err
	}

	return computer.LIP, nil
}
