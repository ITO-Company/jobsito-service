package request

import (
	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type RequestService interface {
	Create(input RequestCreateDto) (*dto.RequestResponseDto, error)
	FindById(id string) (*dto.RequestResponseDto, error)
	FindByIssueId(issueId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.RequestResponseDto], error)
	Update(id string, input RequestUpdateDto) (*dto.RequestResponseDto, error)
	Review(id string, input RequestReviewDto) (*dto.RequestResponseDto, error)
}

type Service struct {
	repo RequestRepo
}

func NewService(repo RequestRepo) RequestService {
	return &Service{repo: repo}
}

func (s *Service) FindByIssueId(issueId string, opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.RequestResponseDto], error) {
	requests, total, err := s.repo.FindByIssueId(issueId, opts)
	if err != nil {
		return nil, err
	}

	dtos := dto.RequestToListDto(requests)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.RequestResponseDto]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) Create(input RequestCreateDto) (*dto.RequestResponseDto, error) {
	issue, err := s.repo.FindIssueById(input.IssueId)
	if err != nil {
		return nil, err
	}

	var request model.Request
	copier.Copy(&request, input)
	request.ID = uuid.New()
	request.Status = enum.StatusPending
	request.FollowupIssueID = issue.ID
	request.FollowupIssue = *issue

	if err := s.repo.Create(&request); err != nil {
		return nil, err
	}

	responseDto := dto.RequestToResponseDto(&request)
	return &responseDto, nil
}

func (s *Service) FindById(id string) (*dto.RequestResponseDto, error) {
	request, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	responseDto := dto.RequestToResponseDto(request)
	return &responseDto, nil
}

func (s *Service) Update(id string, input RequestUpdateDto) (*dto.RequestResponseDto, error) {
	request, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := copier.Copy(request, input); err != nil {
		return nil, err
	}

	if err := s.repo.Update(request); err != nil {
		return nil, err
	}

	responseDto := dto.RequestToResponseDto(request)
	return &responseDto, nil
}

func (s *Service) Review(id string, input RequestReviewDto) (*dto.RequestResponseDto, error) {
	request, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	request.Status = enum.StatusEnum(input.Status)
	request.CompanyComment = input.CompanyComment

	if err := s.repo.Update(request); err != nil {
		return nil, err
	}

	responseDto := dto.RequestToResponseDto(request)
	return &responseDto, nil
}
