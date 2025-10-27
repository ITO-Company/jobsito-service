package jobposting

import "time"

type JobPostingCreateDto struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Requirement     string    `json:"requirement"`
	SalaryMin       string    `json:"salary_min"`
	SalaryMax       string    `json:"salary_max"`
	WorkType        string    `json:"work_type"`
	ExperienceLevel string    `json:"experience_level"`
	Location        string    `json:"location"`
	IsRemote        bool      `json:"is_remote"`
	IsHibrid        bool      `json:"is_hibrid"`
	ContractType    string    `json:"contract_type"`
	Benefit         string    `json:"benefit"`
	Status          string    `json:"status"`
	ExpiresAt       time.Time `json:"expires_at"`
	Tags            []string  `json:"tags"`
}

type JobPostingUpdateDto struct {
	Title           *string    `json:"title"`
	Description     *string    `json:"description"`
	Requirement     *string    `json:"requirement"`
	SalaryMin       *string    `json:"salary_min"`
	SalaryMax       *string    `json:"salary_max"`
	WorkType        *string    `json:"work_type"`
	ExperienceLevel *string    `json:"experience_level"`
	Location        *string    `json:"location"`
	IsRemote        *bool      `json:"is_remote"`
	IsHibrid        *bool      `json:"is_hibrid"`
	ContractType    *string    `json:"contract_type"`
	Benefit         *string    `json:"benefit"`
	Status          *string    `json:"status"`
	ExpiresAt       *time.Time `json:"expires_at"`
}
