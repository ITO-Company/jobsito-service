package src

import (
	"github.com/ito-company/jobsito-service/config"
	"github.com/ito-company/jobsito-service/src/profile/company"
	jobseeker "github.com/ito-company/jobsito-service/src/profile/job_seeker"
)

type Container struct {
	JobSeekerHandler jobseeker.JobSeekerHandler
	CompanyHandler   company.CompanyHandler
}

func SetupContainer() *Container {
	repo := jobseeker.NewRepo(config.DB)
	service := jobseeker.NewService(repo)
	handler := jobseeker.NewHandler(service)

	companyRepo := company.NewRepo(config.DB)
	companyService := company.NewService(companyRepo)
	companyHandler := company.NewHandler(companyService)

	return &Container{
		JobSeekerHandler: handler,
		CompanyHandler:   companyHandler,
	}
}
