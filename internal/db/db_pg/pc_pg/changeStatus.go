package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Изменяем статус ПК в таблице и возвращаем lip
func ChangeStatus(id, status string) (string, error) {
	db := pggorm.DB
	if db == nil {
		return "", errors.New("DB connection is nil")
	}
	var computer models.Computer
	if err := db.Model(&models.Computer{}).
		Where("id = ?", id).
		Update("status", status).
		Select("l_ip").
		First(&computer).Error; err != nil {
		return "", err
	}
	return computer.LIP, nil
}
