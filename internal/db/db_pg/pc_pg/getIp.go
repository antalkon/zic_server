package pcpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetIpByID(id string) (string, error) {
	db := pggorm.DB
	if db == nil {
		return "", errors.New("DB connection is nil")
	}

	var pc models.Computer
	err := db.Where("id =?", id).First(&pc).Error

	if err != nil {
		return "", err
	}

	return pc.LIP, nil
}
