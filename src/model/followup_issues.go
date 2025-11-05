package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/enum"
	"gorm.io/gorm"
)

type FollowupIssue struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title       string
	Description string
	DueDate     time.Time
	Status      enum.StatusEnum

	FollowupMilestoneId uuid.UUID         `gorm:"not null"`
	FollowupMilestone   FollowupMilestone `gorm:"foreignKey:FollowupMilestoneId;references:ID"`

	Requests []Request `gorm:"foreignKey:FollowupIssueID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
