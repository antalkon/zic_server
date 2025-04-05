package models

import (
	"github.com/google/uuid"
)

type Role struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name  string    `gorm:"size:255;not null"`
	Desc  string    `gorm:"size:255"`
	Users []User    `gorm:"foreignKey:RoleID"`
}
