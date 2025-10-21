package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowupMilestone struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedById        string
	CreatedByType      string
	AssignedToId       string
	AssignedToType     string
	Title              string
	Description        string
	Status             string

	CompanyProfileId  uuid.UUID        `gorm:"not null"`
	CompanyProfile    CompanyProfile `gorm:"foreignKey:CompanyProfileId;references:ID"`

	FollowupIssues []FollowupIssue `gorm:"foreignKey:FollowupMilestoneId;references:ID"`
		
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}