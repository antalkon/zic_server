package models

import "time"

type Class struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Number   uint16 `gorm:"default:0" json:"number"`
	Letter   string `gorm:"size:25" json:"letter"`
	FolderID string `gorm:"unique;not null" json:"folder_id"`
}

type PublicTasks struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Classes   string    `gorm:"not null" json:"classes"` // JSONB для массива строк
	Title     string    `gorm:"not null" json:"title"`
	FolderID  string    `gorm:"default:null" json:"folder_id"`
	File      string    `gorm:"default:null" json:"file"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // autoCreateTime для автоматической даты создания
}

type DeliveredTasks struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement:false" json:"id"` // Используем ID как идентификатор задачи
	ClassID   uint64    `gorm:"not null" json:"class_id"`
	Students  string    `gorm:"not null" json:"students"`
	Coment    string    `gorm:"default:'none'" json:"comment"`
	File      string    `gorm:"default:'none'" json:"file"`
	Checked   bool      `gorm:"default:false" json:"checked"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
