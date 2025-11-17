package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type MilestoneResponseDto struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func MilestoneToResponse(m *model.FollowupMilestone) MilestoneResponseDto {
	var dto MilestoneResponseDto
	copier.Copy(&dto, m)

	dto.ID = m.ID.String()

	return dto
}

func MilestoneToListDto(m []model.FollowupMilestone) []MilestoneResponseDto {
	out := make([]MilestoneResponseDto, len(m))
	for i := range m {
		out[i] = MilestoneToResponse(&m[i])
	}
	return out
}
