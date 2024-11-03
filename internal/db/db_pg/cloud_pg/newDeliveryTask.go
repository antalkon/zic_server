package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func SaveDeliveryTasl(dt *models.DeliveredTasks) error {
	db := pggorm.DB
	if db == nil {
		return errors.New("Error creating database SaveDeliveryTasks")
	}

	db.Create(dt)
	return nil
}
