package models

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"size:255;not null"`
	Number      int
	Description string     `gorm:"size:255"`
	ImageID     uuid.UUID  `gorm:"type:uuid"`
	Status      string     `gorm:"size:255;not null"` // active | blocked
	CreatedAt   time.Time  `gorm:"not null"`
	Computers   []Computer `gorm:"foreignKey:Room"`
}
