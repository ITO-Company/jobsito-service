package application

import (
	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ApplicationService interface {
	Create(input *ApplicationCreateDto, jobSeekerID string) (*dto.ApplicationResponse, error)
	Update(input *ApplicationUpdateDto, applicationID string) (*dto.ApplicationResponse, error)
	FindByID(id string) (*dto.ApplicationResponse, error)
	FindAllByJobSeeker(opts *helper.FindAllOptions, jobSeekerID string) ([]dto.ApplicationResponse, error)
	FindAllByJobPostingAndCompany(opts *helper.FindAllOptions, jobPostingID string, companyID string) ([]dto.ApplicationResponse, error)
}

type Service struct {
	repo ApplicationRepo
}

func NewApplicationService(repo ApplicationRepo) ApplicationService {
	return &Service{repo: repo}
}

func (s *Service) Create(input *ApplicationCreateDto, jobSeekerID string) (*dto.ApplicationResponse, error) {
	job, err := s.repo.FindJobPostingById(input.JobPostingID)
	if err != nil {
		return nil, err
	}

	var apply model.Application

	copier.Copy(&apply, input)
	apply.ID = uuid.New()
	apply.JobSeekerId, _ = uuid.Parse(jobSeekerID)
	apply.JobPostingId = job.ID
	if err := s.repo.Create(&apply); err != nil {
		return nil, err
	}

	dto := dto.ApplicationToDto(&apply)

	return &dto, nil
}

func (s *Service) Update(input *ApplicationUpdateDto, applicationID string) (*dto.ApplicationResponse, error) {
	application, err := s.repo.FindById(applicationID)
	if err != nil {
		return nil, err
	}

	returnValue := &dto.ApplicationResponse{}

	err = s.repo.(*Repo).db.Transaction(func(tx *gorm.DB) error {
		copier.Copy(application, input)
		if err := tx.Save(application).Error; err != nil {
			return err
		}

		if input.IsAccepted != nil && *input.IsAccepted {
			var jobPosting model.JobPosting
			if err := tx.Where("id = ?", application.JobPostingId).First(&jobPosting).Error; err != nil {
				return err
			}
			jobPosting.IsClosed = true
			if err := tx.Save(&jobPosting).Error; err != nil {
				return err
			}
		}

		*returnValue = dto.ApplicationToDto(application)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return returnValue, nil
}

func (s *Service) FindByID(id string) (*dto.ApplicationResponse, error) {
	application, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	dto := dto.ApplicationToDto(application)

	return &dto, nil
}

func (s *Service) FindAllByJobSeeker(opts *helper.FindAllOptions, jobSeekerID string) ([]dto.ApplicationResponse, error) {
	applications, _, err := s.repo.FindAllByJobSeeker(jobSeekerID, opts)
	if err != nil {
		return nil, err
	}

	var dtos []dto.ApplicationResponse
	for _, application := range applications {
		dtos = append(dtos, dto.ApplicationToDto(&application))
	}

	return dtos, nil
}

func (s *Service) FindAllByJobPostingAndCompany(opts *helper.FindAllOptions, jobPostingID string, companyID string) ([]dto.ApplicationResponse, error) {
	applications, _, err := s.repo.FindAllByJobPostingAndCompany(jobPostingID, companyID, opts)
	if err != nil {
		return nil, err
	}

	var dtos []dto.ApplicationResponse
	for _, application := range applications {
		dtos = append(dtos, dto.ApplicationToDto(&application))
	}

	return dtos, nil
}
