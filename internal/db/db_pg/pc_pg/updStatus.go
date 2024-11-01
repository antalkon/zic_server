package pcpg

import (
	"fmt"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Изменяем статус ПК
func UpdateComputerStatus(ip string, status bool) error {
	db := pggorm.DB
	if db == nil {
		return fmt.Errorf("DB connection is not established")
	}
	if err := db.Model(&models.Computer{}).Where("l_ip = ?", ip).Update("on", status).Error; err != nil {
		return fmt.Errorf("failed to update status for IP %s: %w", ip, err)
	}
	return nil
}
