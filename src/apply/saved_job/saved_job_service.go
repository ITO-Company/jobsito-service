package savedjob

import (
	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/model"
)

type SavedJobService interface {
	Create(input *SavedJobCreateDto, jobSeekerID string) (*dto.SavedJobResponse, error)
	Delete(id string, jobSeekerID string) error
	FindAllByJobSeeker(opts *helper.FindAllOptions, jobSeekerID string) (*helper.PaginatedResponse[dto.SavedJobResponse], error)
}

type Service struct {
	repo SavedJobRepo
}

func NewService(repo SavedJobRepo) SavedJobService {
	return &Service{repo: repo}
}

func (s *Service) Create(input *SavedJobCreateDto, jobSeekerID string) (*dto.SavedJobResponse, error) {
	jobPosting, err := s.repo.FindJobPostingById(input.JobPostingID)
	if err != nil {
		return nil, err
	}

	// Check if already saved
	_, err = s.repo.FindByJobSeekerAndJobPosting(jobSeekerID, input.JobPostingID)
	if err == nil {
		// Already exists
		return nil, nil
	}

	savedJob := &model.SavedJob{
		ID:                 uuid.New(),
		JobSeekerProfileID: uuid.MustParse(jobSeekerID),
		JobPostingId:       jobPosting.ID,
		JobPosting:         *jobPosting,
	}

	if err := s.repo.Create(savedJob); err != nil {
		return nil, err
	}

	response := dto.SavedJobToResponse(savedJob)
	return &response, nil
}

func (s *Service) Delete(id string, jobSeekerID string) error {
	savedJob, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	if savedJob.JobSeekerProfileID.String() != jobSeekerID {
		return nil
	}

	return s.repo.Delete(id)
}

func (s *Service) FindAllByJobSeeker(opts *helper.FindAllOptions, jobSeekerID string) (*helper.PaginatedResponse[dto.SavedJobResponse], error) {
	savedJobs, total, err := s.repo.FindAllByJobSeeker(jobSeekerID, opts)
	if err != nil {
		return nil, err
	}

	responses := dto.SavedJobToListResponse(savedJobs)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.SavedJobResponse]{
		Data:   responses,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}
