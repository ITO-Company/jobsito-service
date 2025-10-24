package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyProfile struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CompanyName string    `gorm:"not null"`
	Password    string    `gorm:"not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Description string
	Website     string
	Phone       string
	Address     string
	Industry    string
	CompanySize string
	LogoUrl     string
	IsVerified  bool

	JobPostings          []JobPosting           `gorm:"foreignKey:CompanyProfileId;references:ID"`
	Interships           []Intership            `gorm:"foreignKey:CompanyProfileId;references:ID"`
	FollowupMilestones   []FollowupMilestone    `gorm:"foreignKey:CompanyProfileId;references:ID"`
	WeeklyCompanyMetrics []WeeklyCompanyMetrics `gorm:"foreignKey:CompanyProfileId;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
