// Computer модель
package models

type Computer struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:false" json:"id"` // Указываем, что ID хранится как bigint
	RoomID uint64 `gorm:"not null" json:"room_id"`                  // Внешний ключ (RoomID) с типом bigint
	Room   Room   `gorm:"foreignKey:RoomID" json:"-"`               // Объект для связи с Room, не сериализуемый в JSON
	Name   string `gorm:"size:55;not null" json:"name"`
	LIP    string `gorm:"size:255;not null;unique" json:"lip"`
	PIP    string `gorm:"size:255;not null;unique" json:"pip"`
	Status string `gorm:"size:255;default:'Н/Д...'" json:"status"`
	On     bool   `gorm:"default:true" json:"on"`
}

type PcCount struct {
	Count int64  `json:"count"`
	On    int64  `json:"on"`
	Color string `json:"color"`
}
