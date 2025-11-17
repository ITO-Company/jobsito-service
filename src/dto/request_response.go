package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type RequestResponseDto struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Status          string    `json:"status"`
	CompanyComment  string    `json:"company_comment"`
	FollowupIssueID string    `json:"followup_issue_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func RequestToResponseDto(m *model.Request) RequestResponseDto {
	var dto RequestResponseDto
	copier.Copy(&dto, m)

	dto.ID = m.ID.String()
	dto.FollowupIssueID = m.FollowupIssueID.String()
	dto.Status = string(m.Status)

	return dto
}

func RequestToListDto(m []model.Request) []RequestResponseDto {
	out := make([]RequestResponseDto, len(m))
	for i := range m {
		out[i] = RequestToResponseDto(&m[i])
	}
	return out
}
