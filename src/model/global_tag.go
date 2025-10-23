package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GlobalTag struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name        string
	Category   string
	Color      string
	IsApproved string
	UsageCount    string

	JobSeekerTags   []JobSeekerTags   `gorm:"foreignKey:GlobalTagID;references:ID"`
	JobPostingTags  []JobPostingTags  `gorm:"foreignKey:GlobalTagID;references:ID"`
	
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}