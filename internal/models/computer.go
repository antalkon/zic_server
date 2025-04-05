package models

import (
	"time"

	"github.com/google/uuid"
)

type Computer struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Room          uuid.UUID `gorm:"type:uuid;not null"`
	Name          string    `gorm:"size:255;not null"`
	Location      string    `gorm:"size:255"`
	Description   string    `gorm:"size:255"`
	PublicIP      string    `gorm:"size:255;not null"`
	LocalIP       string    `gorm:"size:255"`
	OS            string    `gorm:"size:255;not null"`
	ClientVersion string    `gorm:"size:255;not null"`
	Secret        string    `gorm:"size:255;not null"`
	Status        string    `gorm:"size:255;not null"` // on | off
	Blocked       bool      `gorm:"not null;default:false"`
	Bage          string    `gorm:"size:255;not null;default:'none'"`
	CreatedAt     time.Time `gorm:"not null"`
	LastActivity  time.Time `gorm:"not null"`
	UpdatedAt     time.Time
	Comment       string `gorm:"size:255"`
}
