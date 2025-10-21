package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowupIssue struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title                 string
	Description           string
	DueDate               string
	ProgressPercentage    string


	FollowupMilestoneId uuid.UUID  `gorm:"not null"`
	FollowupMilestone   FollowupMilestone `gorm:"foreignKey:FollowupMilestoneId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}