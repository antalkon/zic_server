package cloudpg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

func GetDeliveryByClassId(classId string) ([]models.DeliveredTasks, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("Error creating database GetDeliveryByRoomId")
	}
	var deliveredTasks []models.DeliveredTasks
	if err := db.Where("class_id =?", classId).Find(&deliveredTasks).Error; err != nil {
		return nil, err
	}
	return deliveredTasks, nil
}
