package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobSeekerTags struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;"`
	ProficiencyLevel string

	JobSeekerProfileID uuid.UUID        `gorm:"not null"`
	JobSeekerProfile   JobSeekerProfile `gorm:"foreignKey:JobSeekerProfileID;references:ID"`

	GlobalTagID uuid.UUID `gorm:"type:uuid;not null"`
	GlobalTag   GlobalTag `gorm:"foreignKey:GlobalTagID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
