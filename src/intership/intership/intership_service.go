package intership

import (
	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type IntershipService interface {
	Create(companyId string, input IntershipCreateDto) (*dto.IntershipResponseDto, error)
	FindById(id string) (*dto.IntershipResponseDto, error)
	FindAll(companyID string, jobSeekerID string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IntershipResponseDto], error)
	FindByIdWithOverview(id string) (*dto.IntershipOverviewDto, error)
	FindAllWithOverview(companyID string, jobSeekerID string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IntershipOverviewDto], error)
	FindByIdWithDetails(id string) (*model.Intership, error)
}

type Service struct {
	repo IntershipRepo
}

func NewService(repo IntershipRepo) IntershipService {
	return &Service{repo: repo}
}

func (s *Service) Create(companyId string, input IntershipCreateDto) (*dto.IntershipResponseDto, error) {
	company, err := s.repo.FindCompanyById(companyId)
	if err != nil {
		return nil, err
	}

	jobSeeker, err := s.repo.FindJobSeekerById(input.JobSeekerProfileID)
	if err != nil {
		return nil, err
	}

	jobPosting, err := s.repo.FindJobPostingById(input.JobPostingId)
	if err != nil {
		return nil, err
	}

	var intership model.Intership
	copier.Copy(&intership, input)
	intership.ID = uuid.New()
	intership.Status = enum.StatusActive
	intership.CompanyProfileId = company.ID
	intership.CompanyProfile = *company
	intership.JobSeekerProfileID = jobSeeker.ID
	intership.JobSeekerProfile = *jobSeeker
	intership.JobPostingId = jobPosting.ID
	intership.JobPosting = *jobPosting

	if err := s.repo.Create(intership); err != nil {
		return nil, err
	}

	response := dto.IntershipToResponse(&intership)
	return &response, nil
}

func (s *Service) FindById(id string) (*dto.IntershipResponseDto, error) {
	intership, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := dto.IntershipToResponse(intership)
	return &response, nil
}

func (s *Service) FindAll(companyID string, jobSeekerID string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IntershipResponseDto], error) {
	interships, total, err := s.repo.FindAll(companyID, jobSeekerID, opts)
	if err != nil {
		return nil, err
	}

	dtos := dto.IntershipToListDto(interships)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.IntershipResponseDto]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindByIdWithOverview(id string) (*dto.IntershipOverviewDto, error) {
	intership, err := s.repo.FindByIdWithDetails(id)
	if err != nil {
		return nil, err
	}

	response := dto.IntershipToOverview(intership)
	return &response, nil
}

func (s *Service) FindAllWithOverview(companyID string, jobSeekerID string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IntershipOverviewDto], error) {
	interships, total, err := s.repo.FindAllWithDetails(companyID, jobSeekerID, opts)
	if err != nil {
		return nil, err
	}

	dtos := dto.IntershipToOverviewList(interships)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.IntershipOverviewDto]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindByIdWithDetails(id string) (*model.Intership, error) {
	intership, err := s.repo.FindByIdWithDetails(id)
	if err != nil {
		return nil, err
	}

	return intership, nil
}
