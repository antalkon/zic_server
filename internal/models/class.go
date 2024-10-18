// Room модель
package models

type Room struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement:false" json:"id"` // Указываем, что ID хранится как bigint
	Name      string     `gorm:"size:25;not null;unique" json:"name"`
	On        bool       `gorm:"default:true" json:"on"`
	Computers []Computer `gorm:"foreignKey:RoomID" json:"computers"` // Указываем внешний ключ для связи
}

type RoomCount struct {
	Count int64  `json:"count"`
	On    int64  `json:"on"`
	Color string `json:"color"`
}
