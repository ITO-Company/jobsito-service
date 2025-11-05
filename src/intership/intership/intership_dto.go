package intership

import "time"

type IntershipCreateDto struct {
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	JobPostingId       string    `json:"job_posting_id"`
	JobSeekerProfileID string    `json:"job_seeker_profile_id"`
}

type IntershipUpdateDto struct {
	StartDate          *time.Time `json:"start_date"`
	EndDate            *time.Time `json:"end_date"`
	Status             *string    `json:"status"`
	JobPostingId       *string    `json:"job_posting_id"`
	JobSeekerProfileID *string    `json:"job_seeker_profile_id"`
	CompanyProfileId   *string    `json:"company_profile_id"`
}
