package models

import "time"

// б модель cpu
type CPULoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	Usage     float64   `gorm:"type:numeric(5,2);not null"`
}

// б модель ram
type RAMLoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	Usage     float64   `gorm:"type:numeric(5,2);not null"` // Процент использования ОЗУ
}

// б модель network
type NetworkLoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	BytesRecv uint64    `gorm:"not null"` // Количество входящего трафика
	BytesSent uint64    `gorm:"not null"` // Количество исходящего трафика
}
