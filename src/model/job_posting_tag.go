package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobPostingTags struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`

	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	GlobalTagID uuid.UUID `gorm:"type:uuid;not null"`
	GlobalTag   GlobalTag `gorm:"foreignKey:GlobalTagID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
