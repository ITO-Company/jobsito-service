package src

import (
	"github.com/ito-company/jobsito-service/config"
	"github.com/ito-company/jobsito-service/src/apply/application"
	savedjob "github.com/ito-company/jobsito-service/src/apply/saved_job"
	"github.com/ito-company/jobsito-service/src/intership/intership"
	"github.com/ito-company/jobsito-service/src/intership/issue"
	"github.com/ito-company/jobsito-service/src/intership/milestone"
	"github.com/ito-company/jobsito-service/src/intership/request"
	"github.com/ito-company/jobsito-service/src/kpi"
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
	MilestoneHandler   milestone.MilestoneHandler
	IssueHandler       issue.IssueHandler
	RequestHandler     request.RequestHandler
	SavedJobHandler    savedjob.SavedJobHandler
	KPIHandler         kpi.KPIHandler
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

	milestoneRepo := milestone.NewRepo(config.DB)
	milestoneService := milestone.NewService(milestoneRepo)
	milestoneHandler := milestone.NewHandler(milestoneService)

	issueRepo := issue.NewRepo(config.DB)
	issueService := issue.NewService(issueRepo)
	issueHandler := issue.NewHandler(issueService)

	requestRepo := request.NewRepo(config.DB)
	requestService := request.NewService(requestRepo)
	requestHandler := request.NewHandler(requestService)

	savedJobRepo := savedjob.NewRepo(config.DB)
	savedJobService := savedjob.NewService(savedJobRepo)
	savedJobHandler := savedjob.NewHandler(savedJobService)

	// KPI initialization
	milestoneKPIRepo := kpi.NewRepo(config.DB)
	milestoneKPIService := kpi.NewMilestoneKPIService(milestoneKPIRepo)

	issueKPIRepo := kpi.NewIssueKPIRepo(config.DB)
	issueKPIService := kpi.NewIssueKPIService(issueKPIRepo)

	requestKPIRepo := kpi.NewRequestKPIRepo(config.DB)
	requestKPIService := kpi.NewRequestKPIService(requestKPIRepo)

	conversionKPIRepo := kpi.NewConversionKPIRepo(config.DB)
	conversionKPIService := kpi.NewConversionKPIService(conversionKPIRepo)

	kpiHandler := kpi.NewHandler(milestoneKPIService, issueKPIService, requestKPIService, conversionKPIService)

	return &Container{
		JobSeekerHandler:   handler,
		CompanyHandler:     companyHandler,
		GlobalTagHandler:   globalTagHandler,
		IssueHandler:       issueHandler,
		MilestoneHandler:   milestoneHandler,
		JobPostingHandler:  jobpostingHandler,
		IntershipHandler:   intershipHandler,
		ApplicationHandler: applicationHandler,
		RequestHandler:     requestHandler,
		SavedJobHandler:    savedJobHandler,
		KPIHandler:         kpiHandler,
	}
}
