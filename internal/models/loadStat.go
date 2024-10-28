package models

import "time"

type CPULoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	Usage     float64   `gorm:"type:numeric(5,2);not null"`
}

type RAMLoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	Usage     float64   `gorm:"type:numeric(5,2);not null"` // Процент использования ОЗУ
}

type NetworkLoad struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `gorm:"not null"`
	BytesRecv uint64    `gorm:"not null"` // Количество входящего трафика
	BytesSent uint64    `gorm:"not null"` // Количество исходящего трафика
}
