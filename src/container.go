package src

import (
	"github.com/ito-company/jobsito-service/config"
	"github.com/ito-company/jobsito-service/src/profile/company"
	globaltags "github.com/ito-company/jobsito-service/src/profile/global_tags"
	jobseeker "github.com/ito-company/jobsito-service/src/profile/job_seeker"
)

type Container struct {
	JobSeekerHandler jobseeker.JobSeekerHandler
	CompanyHandler   company.CompanyHandler
	GlobalTagHandler globaltags.GlobalTagHandler
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

	return &Container{
		JobSeekerHandler: handler,
		CompanyHandler:   companyHandler,
		GlobalTagHandler: globalTagHandler,
	}
}
