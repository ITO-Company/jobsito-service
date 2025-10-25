package globaltags

import (
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
)

type GlobalTagsService interface {
	FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.GlobalTagResponse], error)
}

type Service struct {
	repo GlobalTagsRepo
}

func NewService(repo GlobalTagsRepo) GlobalTagsService {
	return &Service{repo: repo}
}

func (s *Service) FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.GlobalTagResponse], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := dto.GlobalTagToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.GlobalTagResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}
