package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationStatusHistory struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	PreviousStatus string
	NewStatus      string
	Notes          string
	ChangedById    string
	ChangeByType   string

	ApplicationId uuid.UUID   `gorm:"not null"`
	Application   Application `gorm:"foreignKey:ApplicationId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
