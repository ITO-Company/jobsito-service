package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobView struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	ViewerId         string
	ViewerType       string
	DurationSeconds  string
	Source           string
	Metadata         string

	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}