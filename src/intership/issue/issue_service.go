package issue

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type IssueService interface {
	Create(input IssueCreateDto) (*dto.IssueResponseDto, error)
	FindById(id string) (*dto.IssueResponseDto, error)
	FindAll(milestoneId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IssueResponseDto], error)
	Update(id string, input IssueUpdateDto) (*dto.IssueResponseDto, error)
}

type Service struct {
	repo IssueRepo
}

func NewService(repo IssueRepo) IssueService {
	return &Service{repo: repo}
}

func (s *Service) FindAll(milestoneId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.IssueResponseDto], error) {
	issues, total, err := s.repo.FindAll(milestoneId, opts)
	if err != nil {
		return nil, err
	}

	dtos := dto.IssueToListDto(issues)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.IssueResponseDto]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) Create(input IssueCreateDto) (*dto.IssueResponseDto, error) {
	milestone, err := s.repo.FindMilestoneById(input.MilestoneId)
	if err != nil {
		return nil, err
	}

	var issue model.FollowupIssue
	copier.Copy(&issue, input)
	issue.FollowupMilestoneId = milestone.ID
	issue.FollowupMilestone = *milestone

	if err := s.repo.Create(&issue); err != nil {
		return nil, err
	}

	responseDto := dto.IssueToResponseDto(&issue)
	return &responseDto, nil
}

func (s *Service) FindById(id string) (*dto.IssueResponseDto, error) {
	issue, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	responseDto := dto.IssueToResponseDto(issue)
	return &responseDto, nil
}

func (s *Service) Update(id string, input IssueUpdateDto) (*dto.IssueResponseDto, error) {
	issue, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(issue, input); err != nil {
		return nil, err
	}

	if err := s.repo.Update(issue); err != nil {
		return nil, err
	}

	responseDto := dto.IssueToResponseDto(issue)
	return &responseDto, nil
}
