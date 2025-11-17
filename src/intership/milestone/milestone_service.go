package milestone

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type MilestoneService interface {
	Create(input MilestoneCreateDto, company_id string) (*dto.MilestoneResponseDto, error)
	FindById(id string) (*dto.MilestoneResponseDto, error)
	FindAll(intershipId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.MilestoneResponseDto], error)
	Update(id string, input MilestoneUpdateDto, userId string) (*dto.MilestoneResponseDto, error)
}

type Service struct {
	repo MilestoneRepo
}

func NewService(repo MilestoneRepo) MilestoneService {
	return &Service{repo: repo}
}

func (s *Service) FindAll(intershipId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.MilestoneResponseDto], error) {
	milestones, total, err := s.repo.FindAll(intershipId, opts)
	if err != nil {
		return nil, err
	}

	dtos := dto.MilestoneToListDto(milestones)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.MilestoneResponseDto]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) Create(input MilestoneCreateDto, company_id string) (*dto.MilestoneResponseDto, error) {
	intership, err := s.repo.FindIntershipById(input.IntershipId)
	if err != nil {
		return nil, err
	}

	company, err := s.repo.FindCompanyById(company_id)
	if err != nil {
		return nil, err
	}

	var milestone model.FollowupMilestone
	copier.Copy(&milestone, input)
	milestone.ID = uuid.New()
	milestone.IntershipId = intership.ID
	milestone.Intership = *intership
	milestone.Status = enum.StatusPending
	milestone.CompanyProfileId = company.ID
	milestone.CompanyProfile = *company

	if err := s.repo.Create(&milestone); err != nil {
		return nil, err
	}

	response := dto.MilestoneToResponse(&milestone)
	return &response, nil
}

func (s *Service) FindById(id string) (*dto.MilestoneResponseDto, error) {
	milestone, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	response := dto.MilestoneToResponse(milestone)
	return &response, nil
}

func (s *Service) Update(id string, input MilestoneUpdateDto, userId string) (*dto.MilestoneResponseDto, error) {
	milestone, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	company, err := s.repo.FindCompanyById(userId)
	if err != nil {
		return nil, err
	}

	if milestone.CompanyProfileId != company.ID {
		return nil, fmt.Errorf("can't edit another company milestone")
	}

	if err := copier.Copy(milestone, input); err != nil {
		return nil, err
	}

	if err := s.repo.Update(milestone); err != nil {
		return nil, err
	}

	response := dto.MilestoneToResponse(milestone)
	return &response, nil
}
