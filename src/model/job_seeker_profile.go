package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobSeekerProfile struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey;"`	
	Name                string
	Bio                 string
	Phone               string
	Location            string
	CvUrl               string
	PortfolioUrl        string
	ExpectedSalaryMin   string
	ExpectedSalaryMax   string
	Availability        string
	Skills              string
	Experience          string
	IsActive            bool

	Application       []Application `gorm:"foreignKey:JobSeekerProfileID;references:ID"`
	JobSeekerTags    []JobSeekerTags `gorm:"foreignKey:JobSeekerProfileID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}