package dto

import (
	"time"

	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type SavedJobResponse struct {
	ID         string             `json:"id"`
	JobPosting JobPostingResponse `json:"job_posting"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func SavedJobToResponse(m *model.SavedJob) SavedJobResponse {
	var dto SavedJobResponse
	copier.Copy(&dto, m)

	dto.ID = m.ID.String()
	dto.JobPosting = JobPostingToDto(&m.JobPosting)

	return dto
}

func SavedJobToListResponse(m []model.SavedJob) []SavedJobResponse {
	out := make([]SavedJobResponse, len(m))
	for i := range m {
		out[i] = SavedJobToResponse(&m[i])
	}
	return out
}
