package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/enum"
	"gorm.io/gorm"
)

type Request struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title          string
	Description    string
	Status         enum.StatusEnum `gorm:"type:varchar(20)"`
	CompanyComment string          // Comentario de la empresa al aprobar/rechazar

	FollowupIssueID uuid.UUID     `gorm:"not null"`
	FollowupIssue   FollowupIssue `gorm:"foreignKey:FollowupIssueID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
