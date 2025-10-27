package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobPosting struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title           string
	Description     string
	Requirement     string
	SalaryMin       string
	SalaryMax       string
	WorkType        string
	ExperienceLevel string
	Location        string
	IsRemote        bool
	IsHibrid        bool
	ContractType    string
	Benefit         string
	Status          string
	ExpiresAt       time.Time

	CompanyProfileId uuid.UUID      `gorm:"not null"`
	CompanyProfile   CompanyProfile `gorm:"foreignKey:CompanyProfileId;references:ID"`

	Applications   []Application    `gorm:"foreignKey:JobPostingId;references:ID"`
	DailyJobStats  []DailyJobStats  `gorm:"foreignKey:JobPostingId;references:ID"`
	JobView        []JobView        `gorm:"foreignKey:JobPostingId;references:ID"`
	Intership      []Intership      `gorm:"foreignKey:JobPostingId;references:ID"`
	JobPostingTags []JobPostingTags `gorm:"foreignKey:JobPostingId;references:ID"`
	SavedJob       []SavedJob       `gorm:"foreignKey:JobPostingId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
