package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type IssueResponseDto struct {
	ID                  string    `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	DueDate             time.Time `json:"due_date"`
	Status              string    `json:"status"`
	FollowupMilestoneId string    `json:"followup_milestone_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

func IssueToResponseDto(m *model.FollowupIssue) IssueResponseDto {
	var dto IssueResponseDto
	copier.Copy(&dto, m)

	dto.ID = m.ID.String()

	return dto
}

func IssueToListDto(m []model.FollowupIssue) []IssueResponseDto {
	out := make([]IssueResponseDto, len(m))
	for i := range m {
		out[i] = IssueToResponseDto(&m[i])
	}
	return out
}
