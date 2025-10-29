package application

type ApplicationCreateDto struct {
	JobPostingID   string `json:"job_posting_id"`
	CoverLetter    string `json:"cover_letter"`
	ProposedSalary string `json:"proposed_salary"`
}

type ApplicationUpdateDto struct {
	IsAccepted *bool `json:"is_accepted"`
}
