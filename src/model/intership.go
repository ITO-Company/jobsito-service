package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/enum"
	"gorm.io/gorm"
)

type Intership struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	StartDate time.Time
	EndDate   time.Time
	Status    enum.StatusEnum

	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	JobSeekerProfileID uuid.UUID        `gorm:"not null"`
	JobSeekerProfile   JobSeekerProfile `gorm:"foreignKey:JobSeekerProfileID;references:ID"`

	CompanyProfileId uuid.UUID      `gorm:"not null"`
	CompanyProfile   CompanyProfile `gorm:"foreignKey:CompanyProfileId;references:ID"`

	Milestones []FollowupMilestone `gorm:"foreignKey:IntershipId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
