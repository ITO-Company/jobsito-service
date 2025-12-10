package jobposting

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type JobPostingService interface {
	Create(companyId string, input JobPostingCreateDto) (*dto.JobPostingResponse, error)
	Update(companyId string, jobId string, input JobPostingUpdateDto) (*dto.JobPostingResponse, error)
	SoftDelete(id string) error
	AddTagToJobPosting(jobPostingId, tagId string) error
	RemoveTagFromJobPosting(jobPostingId, tagId string) error
	AuthorizeCompanyAction(companyId string, jobPostingId string) error
	FindAll(opts *helper.FindAllOptions, tagIDs []string, companyID string) (*helper.PaginatedResponse[dto.JobPostingResponse], error)
	FindById(id string) (*dto.JobPostingResponse, error)
	FindAllWithApplications(companyID string) ([]model.JobPosting, error)
	FindByIdWithApplications(id string) (*model.JobPosting, error)
}

type Service struct {
	repo JobPostingRepo
}

func NewService(repo JobPostingRepo) JobPostingService {
	return &Service{repo: repo}
}

func (s *Service) Create(companyId string, input JobPostingCreateDto) (*dto.JobPostingResponse, error) {
	company, err := s.repo.FindCompanyById(companyId)
	if err != nil {
		return nil, err
	}

	var job model.JobPosting
	copier.Copy(&job, &input)
	job.ID = uuid.New()
	job.CompanyProfileId = company.ID
	job.CompanyProfile = *company
	job.IsClosed = false

	err = s.repo.CreateWithTags(job, input.Tags)
	if err != nil {
		return nil, err
	}

	response := dto.JobPostingToDto(&job)
	return &response, nil
}

func (s *Service) Update(companyId string, jobId string, input JobPostingUpdateDto) (*dto.JobPostingResponse, error) {
	job, err := s.repo.FindById(jobId)
	if err != nil {
		return nil, err
	}

	if job.CompanyProfileId.String() != companyId {
		return nil, fmt.Errorf("unauthorized to update this job posting")
	}

	opt := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	}

	if err := copier.CopyWithOption(job, &input, opt); err != nil {
		return nil, err
	}

	err = s.repo.Update(*job)
	if err != nil {
		return nil, err
	}

	response := dto.JobPostingToDto(job)
	return &response, nil
}

func (s *Service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *Service) AddTagToJobPosting(jobPostingId, tagId string) error {
	return s.repo.AddTagToJobPosting(jobPostingId, tagId)
}

func (s *Service) RemoveTagFromJobPosting(jobPostingId, tagId string) error {
	return s.repo.RemoveTagFromJobPosting(jobPostingId, tagId)
}

func (s *Service) AuthorizeCompanyAction(companyId string, jobPostingId string) error {
	job, err := s.repo.FindById(jobPostingId)
	if err != nil {
		return err
	}

	if job.CompanyProfileId.String() != companyId {
		return fmt.Errorf("unauthorized to perform this action on the job posting")
	}

	return nil
}

func (s *Service) FindAll(opts *helper.FindAllOptions, tagIDs []string, companyID string) (*helper.PaginatedResponse[dto.JobPostingResponse], error) {
	finded, total, err := s.repo.FindAll(opts, tagIDs, companyID)
	if err != nil {
		return nil, err
	}
	dtos := dto.JobPostingToListDto(finded)
	
	// Cap the total to the limit (30 max records)
	if total > int64(opts.Limit) {
		total = int64(opts.Limit)
	}
	
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.JobPostingResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindById(id string) (*dto.JobPostingResponse, error) {
	job, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := dto.JobPostingToDto(job)
	return &response, nil
}

func (s *Service) FindAllWithApplications(companyID string) ([]model.JobPosting, error) {
	return s.repo.FindAllWithApplications(companyID)
}

func (s *Service) FindByIdWithApplications(id string) (*model.JobPosting, error) {
	return s.repo.FindByIdWithApplications(id)
}
