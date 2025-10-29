package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type ApplicationResponse struct {
	ID               string            `json:"id"`
	CoverLetter      string            `json:"cover_letter"`
	Status           string            `json:"status"`
	ApplicationNotes string            `json:"Application_notes"`
	ProposedSalary   string            `json:"proposed_salary"`
	StatusChangedAt  string            `json:"status_changed_at"`
	JobSeekerId      string            `json:"job_seeker_id"`
	JobPostingId     string            `json:"job_posting_id"`
	IsAccepted       bool              `json:"is_accepted"`
	JobSeeker        JobSeekerResponse `json:"job_seeker"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

func ApplicationToDto(m *model.Application) ApplicationResponse {
	var dto ApplicationResponse
	copier.Copy(&dto, m)

	if m.JobSeeker.ID != uuid.Nil {
		dto.JobSeeker = JobSeekerToDto(&m.JobSeeker)
	}

	return dto
}

func ApplicationToListDto(m []model.Application) []ApplicationResponse {
	out := make([]ApplicationResponse, len(m))
	for i := range m {
		out[i] = ApplicationToDto(&m[i])
	}
	return out
}
