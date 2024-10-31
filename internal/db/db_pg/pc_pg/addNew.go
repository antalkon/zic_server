package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
	"github.com/antalkon/zic_server/pkg/logger"
)

func AddNewPc(pc *models.Computer) error {
	db := pggorm.DB
	if db == nil {
		err := errors.New("database connection is not established")
		logger.LogError(err)
		return err
	}

	// Проверка на существование записи (если требуется)
	var existingPC models.Computer
	if err := db.Where("l_ip = ? OR p_ip = ?", pc.LIP, pc.PIP).First(&existingPC).Error; err == nil {
		return errors.New("computer with the same local or public IP already exists")
	}

	// Создание новой записи
	if err := db.Create(pc).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
