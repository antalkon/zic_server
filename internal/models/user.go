package models

import (
	"time"

	"github.com/google/uuid"
)

// @Description Структура, описывающая пользователя

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name         string    `gorm:"size:55;not null" json:"name"`
	Surname      string    `gorm:"size:55;not null" json:"surname"`
	Email        string    `gorm:"size:255;not null" json:"email"`
	Phone        string    `gorm:"size:255" json:"phone,omitempty"`
	PasswordHash string    `gorm:"size:255;not null" json:""`   // не возвращается в JSON
	Password     string    `gorm:"-" json:"password,omitempty"` // не сохраняется в БД
	RoleID       uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	Role         Role      `gorm:"foreignKey:RoleID" json:"role"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
