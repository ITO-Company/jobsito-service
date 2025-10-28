package jobseeker

type JobSeekerUpdateDto struct {
	Name              *string `json:"name"`
	Bio               *string `json:"bio"`
	Phone             *string `json:"phone"`
	Location          *string `json:"location"`
	CvUrl             *string `json:"cv_url"`
	PortfolioUrl      *string `json:"portfolio_url"`
	ExpectedSalaryMin *string `json:"expected_salary_min"`
	ExpectedSalaryMax *string `json:"expected_salary_max"`
	Availability      *string `json:"availability"`
	Skills            *string `json:"skills"`
	Experience        *string `json:"experience"`
}
