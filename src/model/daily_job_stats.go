package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DailyJobStats struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	StatsData      string
	TotalViews     string
	UniqueViews    string
	Applications   string
	Saves          string

	JobPostingId uuid.UUID        `gorm:"not null"`
    JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}