package roompg

import (
	"errors"

	"github.com/antalkon/zic_server/internal/models"
	pggorm "github.com/antalkon/zic_server/pkg/db/pg_gorm"
)

// Получение количества ПК с room_id = room и on = true
func GetCountPCOn(room string) (int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, errors.New("DB connection is nil")
	}
	var count int64
	if err := db.Model(&models.Computer{}).Where("room_id = ? AND \"on\" = ?", room, true).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Получение количества ПК с room_id = room и on = false
func GetCountPCOff(room string) (int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, errors.New("DB connection is nil")
	}
	var count int64
	if err := db.Model(&models.Computer{}).Where("room_id = ? AND \"on\" = ?", room, false).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Получение общего количества ПК с room_id = room
func GetTotalPCCount(room string) (int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, errors.New("DB connection is nil")
	}
	var count int64
	if err := db.Model(&models.Computer{}).Where("room_id = ?", room).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Получение количества ПК с room_id = room и status = "blocked"
func GetCountPCBlocked(room string) (int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, errors.New("DB connection is nil")
	}
	var count int64
	if err := db.Model(&models.Computer{}).Where("room_id = ? AND status = ?", room, "blocked").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Получение количества ПК с room_id = room и status = "work"
func GetCountPCWorking(room string) (int64, error) {
	db := pggorm.DB
	if db == nil {
		return 0, errors.New("DB connection is nil")
	}
	var count int64
	if err := db.Model(&models.Computer{}).Where("room_id = ? AND status = ?", room, "work").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Получение массива всех ПК с room_id = room
func GetAllPCsInRoom(room string) ([]models.Computer, error) {
	db := pggorm.DB
	if db == nil {
		return nil, errors.New("DB connection is nil")
	}
	var computers []models.Computer
	if err := db.Where("room_id = ?", room).Find(&computers).Error; err != nil {
		return nil, err
	}
	return computers, nil
}
