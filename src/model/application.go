package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CoverLetter     string
	Status          string
	CompanyNotes    string
	ProposedSalary  string
	StatusChangedAt string
	IsAccepted      bool

	JobSeekerId uuid.UUID        `gorm:"not null"`
	JobSeeker   JobSeekerProfile `gorm:"foreignKey:JobSeekerId;references:ID"`

	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	ApplicationStatusHistory []ApplicationStatusHistory `gorm:"foreignKey:ApplicationId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
