package src

import (
	"github.com/ito-company/jobsito-service/config"
	jobseeker "github.com/ito-company/jobsito-service/src/profile/job_seeker"
)

type Container struct {
	// JobSeeker
	JobSeekerHandler jobseeker.JobSeekerHandler
}

func SetupContainer() *Container {
	repo := jobseeker.NewRepo(config.DB)
	service := jobseeker.NewService(repo)
	handler := jobseeker.NewHandler(service)

	return &Container{
		JobSeekerHandler: handler,
	}
}
