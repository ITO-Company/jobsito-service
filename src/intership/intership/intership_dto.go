package intership

type IntershipCreateDto struct {
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	JobPostingId       string `json:"job_posting_id"`
	JobSeekerProfileID string `json:"job_seeker_profile_id"`
}

type IntershipUpdateDto struct {
	StartDate          *string `json:"start_date"`
	EndDate            *string `json:"end_date"`
	Status             *string `json:"status"`
	JobPostingId       *string `json:"job_posting_id"`
	JobSeekerProfileID *string `json:"job_seeker_profile_id"`
	CompanyProfileId   *string `json:"company_profile_id"`
}
