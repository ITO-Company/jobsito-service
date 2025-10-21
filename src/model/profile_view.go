package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileView struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;"`
	ViewerType       string
	ProfileId        string
	ProfileType      string
	Source           string
	Metadata         string


	JobSeekerProfileID uuid.UUID        `gorm:"not null"`
	JobSeekerProfile   JobSeekerProfile `gorm:"foreignKey:JobSeekerProfileID;references:ID"`


	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}