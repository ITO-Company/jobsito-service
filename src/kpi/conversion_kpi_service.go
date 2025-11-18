package kpi

// ConversionKPIDto holds all KPI metrics for conversion by company or job posting
type ConversionKPIDto struct {
	ContextID                        string  `json:"context_id"`   // Company ID or JobPosting ID
	ContextType                      string  `json:"context_type"` // "company" or "job_posting"
	TotalApplications                int64   `json:"total_applications"`
	AcceptedApplications             int64   `json:"accepted_applications"`
	InitiatedInterships              int64   `json:"initiated_interships"`
	CompletedInterships              int64   `json:"completed_interships"`
	ApplicationAcceptanceRate        float64 `json:"application_acceptance_rate"`
	ConversionToIntershipRate        float64 `json:"conversion_to_intership_rate"`
	IntershipCompletionRate          float64 `json:"intership_completion_rate"`
	AvgTimeAppToAcceptanceDays       float64 `json:"avg_time_app_to_acceptance_days"`
	AvgTimeAcceptanceToIntershipDays float64 `json:"avg_time_acceptance_to_intership_days"`
	AvgProposedSalary                string  `json:"avg_proposed_salary"`
	AvgOfferedSalaryMin              string  `json:"avg_offered_salary_min"`
	AvgOfferedSalaryMax              string  `json:"avg_offered_salary_max"`
}

// ConversionKPIService handles business logic for conversion KPIs
type ConversionKPIService interface {
	GetCompanyConversionKPI(companyID string) (*ConversionKPIDto, error)
	GetJobPostingConversionKPI(jobPostingID string) (*ConversionKPIDto, error)
}

type ConversionKPIServiceImpl struct {
	repo ConversionKPIRepo
}

func NewConversionKPIService(repo ConversionKPIRepo) ConversionKPIService {
	return &ConversionKPIServiceImpl{repo: repo}
}

func (s *ConversionKPIServiceImpl) GetCompanyConversionKPI(companyID string) (*ConversionKPIDto, error) {
	totalApps, err := s.repo.CountApplicationsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	acceptedApps, err := s.repo.CountAcceptedApplicationsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	initiatedInterships, err := s.repo.CountInitiatedIntershipsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	completedInterships, err := s.repo.CountCompletedIntershipsByCompany(companyID)
	if err != nil {
		return nil, err
	}

	timeAppToAcceptance, err := s.repo.FindAverageTimeApplicationToAcceptanceByCompany(companyID)
	if err != nil {
		return nil, err
	}

	timeAcceptanceToIntership, err := s.repo.FindAverageTimeAcceptanceToIntershipByCompany(companyID)
	if err != nil {
		return nil, err
	}

	proposedAvg, offeredMin, offeredMax, err := s.repo.FindAverageSalaryProposedVsOfferedByCompany(companyID)
	if err != nil {
		return nil, err
	}

	// Calculate rates
	appAcceptanceRate := 0.0
	if totalApps > 0 {
		appAcceptanceRate = float64(acceptedApps) / float64(totalApps) * 100
	}

	conversionToIntershipRate := 0.0
	if acceptedApps > 0 {
		conversionToIntershipRate = float64(initiatedInterships) / float64(acceptedApps) * 100
	}

	intershipCompletionRate := 0.0
	if initiatedInterships > 0 {
		intershipCompletionRate = float64(completedInterships) / float64(initiatedInterships) * 100
	}

	return &ConversionKPIDto{
		ContextID:                        companyID,
		ContextType:                      "company",
		TotalApplications:                totalApps,
		AcceptedApplications:             acceptedApps,
		InitiatedInterships:              initiatedInterships,
		CompletedInterships:              completedInterships,
		ApplicationAcceptanceRate:        appAcceptanceRate,
		ConversionToIntershipRate:        conversionToIntershipRate,
		IntershipCompletionRate:          intershipCompletionRate,
		AvgTimeAppToAcceptanceDays:       timeAppToAcceptance.Hours() / 24,
		AvgTimeAcceptanceToIntershipDays: timeAcceptanceToIntership.Hours() / 24,
		AvgProposedSalary:                proposedAvg,
		AvgOfferedSalaryMin:              offeredMin,
		AvgOfferedSalaryMax:              offeredMax,
	}, nil
}

func (s *ConversionKPIServiceImpl) GetJobPostingConversionKPI(jobPostingID string) (*ConversionKPIDto, error) {
	totalApps, err := s.repo.CountApplicationsByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	acceptedApps, err := s.repo.CountAcceptedApplicationsByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	initiatedInterships, err := s.repo.CountInitiatedIntershipsByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	completedInterships, err := s.repo.CountCompletedIntershipsByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	timeAppToAcceptance, err := s.repo.FindAverageTimeApplicationToAcceptanceByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	timeAcceptanceToIntership, err := s.repo.FindAverageTimeAcceptanceToIntershipByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	proposedAvg, offeredMin, offeredMax, err := s.repo.FindAverageSalaryProposedVsOfferedByJobPosting(jobPostingID)
	if err != nil {
		return nil, err
	}

	// Calculate rates
	appAcceptanceRate := 0.0
	if totalApps > 0 {
		appAcceptanceRate = float64(acceptedApps) / float64(totalApps) * 100
	}

	conversionToIntershipRate := 0.0
	if acceptedApps > 0 {
		conversionToIntershipRate = float64(initiatedInterships) / float64(acceptedApps) * 100
	}

	intershipCompletionRate := 0.0
	if initiatedInterships > 0 {
		intershipCompletionRate = float64(completedInterships) / float64(initiatedInterships) * 100
	}

	return &ConversionKPIDto{
		ContextID:                        jobPostingID,
		ContextType:                      "job_posting",
		TotalApplications:                totalApps,
		AcceptedApplications:             acceptedApps,
		InitiatedInterships:              initiatedInterships,
		CompletedInterships:              completedInterships,
		ApplicationAcceptanceRate:        appAcceptanceRate,
		ConversionToIntershipRate:        conversionToIntershipRate,
		IntershipCompletionRate:          intershipCompletionRate,
		AvgTimeAppToAcceptanceDays:       timeAppToAcceptance.Hours() / 24,
		AvgTimeAcceptanceToIntershipDays: timeAcceptanceToIntership.Hours() / 24,
		AvgProposedSalary:                proposedAvg,
		AvgOfferedSalaryMin:              offeredMin,
		AvgOfferedSalaryMax:              offeredMax,
	}, nil
}
