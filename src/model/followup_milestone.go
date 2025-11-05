package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/enum"
	"gorm.io/gorm"
)

type FollowupMilestone struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title       string
	Description string
	Status      enum.StatusEnum

	IntershipId uuid.UUID `gorm:"not null"`
	Intership   Intership `gorm:"foreignKey:IntershipId;references:ID"`

	FollowupIssues []FollowupIssue `gorm:"foreignKey:FollowupMilestoneId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
