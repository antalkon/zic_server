package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// UpdateLipById обновляет поле lip по ID
func UpdateLipById(id int64, ip string) error {
	// Получаем подключение к базе данных
	db := pggorm.DB
	if db == nil {
		return errors.New("database connection is not established")
	}

	// Обновляем поле lip по ID
	if err := db.Model(&models.Computer{}).Where("id = ?", id).Update("l_ip", ip).Error; err != nil {
		return err // Возвращаем ошибку, если обновление не удалось
	}

	return nil // Успешное завершение
}
