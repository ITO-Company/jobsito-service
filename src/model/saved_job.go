package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SavedJob struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`

	JobSeekerProfileID uuid.UUID        `gorm:"not null"`
	JobSeekerProfile   JobSeekerProfile `gorm:"foreignKey:JobSeekerProfileID;references:ID"`

	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
