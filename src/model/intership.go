package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Intership struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	PositionTitle    string
	StartDate        string
	EndDate          string
	Status           string


	JobPostingId uuid.UUID  `gorm:"not null"`
	JobPosting   JobPosting `gorm:"foreignKey:JobPostingId;references:ID"`

	JobSeekerProfileID uuid.UUID        `gorm:"not null"`
	JobSeekerProfile   JobSeekerProfile `gorm:"foreignKey:JobSeekerProfileID;references:ID"`

	CompanyProfileId  uuid.UUID        `gorm:"not null"`
	CompanyProfile    CompanyProfile `gorm:"foreignKey:CompanyProfileId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}