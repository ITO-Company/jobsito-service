package src

import (
	"github.com/ito-company/jobsito-service/config"
	"github.com/ito-company/jobsito-service/src/apply/application"
	"github.com/ito-company/jobsito-service/src/intership/intership"
	jobposting "github.com/ito-company/jobsito-service/src/offer/job_posting"
	"github.com/ito-company/jobsito-service/src/profile/company"
	globaltags "github.com/ito-company/jobsito-service/src/profile/global_tags"
	jobseeker "github.com/ito-company/jobsito-service/src/profile/job_seeker"
)

type Container struct {
	JobSeekerHandler   jobseeker.JobSeekerHandler
	CompanyHandler     company.CompanyHandler
	GlobalTagHandler   globaltags.GlobalTagHandler
	JobPostingHandler  jobposting.JobPostingHandler
	ApplicationHandler application.ApplicationHandler
	IntershipHandler   intership.IntershipHandler
}

func SetupContainer() *Container {
	repo := jobseeker.NewRepo(config.DB)
	service := jobseeker.NewService(repo)
	handler := jobseeker.NewHandler(service)

	companyRepo := company.NewRepo(config.DB)
	companyService := company.NewService(companyRepo)
	companyHandler := company.NewHandler(companyService)

	globalTagRepo := globaltags.NewRepo(config.DB)
	globalTagService := globaltags.NewService(globalTagRepo)
	globalTagHandler := globaltags.NewHandler(globalTagService)

	jobpostingRepo := jobposting.NewRepo(config.DB)
	jobpostingService := jobposting.NewService(jobpostingRepo)
	jobpostingHandler := jobposting.NewHandler(jobpostingService)

	applicationRepo := application.NewRepo(config.DB)
	applicationService := application.NewService(applicationRepo)
	applicationHandler := application.NewHandler(applicationService)

	intershipRepo := intership.NewRepo(config.DB)
	intershipService := intership.NewService(intershipRepo)
	intershipHandler := intership.NewHandler(intershipService)

	return &Container{
		JobSeekerHandler:   handler,
		CompanyHandler:     companyHandler,
		GlobalTagHandler:   globalTagHandler,
		JobPostingHandler:  jobpostingHandler,
		IntershipHandler:   intershipHandler,
		ApplicationHandler: applicationHandler,
	}
}
