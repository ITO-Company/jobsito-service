package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WeeklyCompanyMetrics struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	WeekStart                  string
	JobsPosted                 string
	TotalApplicationsReceived  uint
	AvgApplicationsPerJob      float64
	HiresMade                  string

	CompanyProfileId  uuid.UUID        `gorm:"not null"`
	CompanyProfile    CompanyProfile `gorm:"foreignKey:CompanyProfileId;references:ID"`


	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}