// Computer модель
package models

// модель компьюетар в базе
type Computer struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:false" json:"id"` // Указываем, что ID хранится как bigint
	RoomID uint64 `gorm:"not null" json:"room"`                     // RoomID больше не связан с таблицей Room
	Name   string `gorm:"size:55;not null" json:"name"`
	LIP    string `gorm:"size:255;not null;unique" json:"lip"`
	PIP    string `gorm:"size:255;not null" json:"pip"`
	Status string `gorm:"size:255;default:'Н/Д...'" json:"status"`
	On     bool   `gorm:"default:true" json:"on"`
}

// модель количества компьютеров в базе
type PcCount struct {
	Count int64  `json:"count"`
	On    int64  `json:"on"`
	Color string `json:"color"`
}
