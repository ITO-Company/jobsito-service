package jobseeker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
)

type JobSeekerService interface {
	Signup(dto dto.SignupDto) (string, error)
	Signin(dto dto.SigninDto) (string, error)
}

type Service struct {
	repo JobSeekerRepo
}

func NewService(repo JobSeekerRepo) JobSeekerService {
	return &Service{repo: repo}
}

func (s *Service) Signup(dto dto.SignupDto) (string, error) {
	existing, err := s.repo.FindByEmail(dto.Email)
	if err == nil && existing.ID != uuid.Nil {
		return "", fmt.Errorf("email already in use")
	}

	if dto.Password != dto.ConfirmPassword {
		return "", fmt.Errorf("passwords do not match")
	}

	hashedPassword, err := helper.HashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	jobSeeker := model.JobSeekerProfile{
		ID:       uuid.New(),
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := s.repo.Create(jobSeeker); err != nil {
		return "", err
	}

	token, err := helper.GenerateJwt(jobSeeker.ID.String(), jobSeeker.Email, string(enum.RoleSeeker))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Signin(dto dto.SigninDto) (string, error) {
	seeker, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckPasswordHash(dto.Password, seeker.Password) {
		return "", fmt.Errorf("incorrect credentials")
	}

	token, err := helper.GenerateJwt(seeker.ID.String(), seeker.Email, string(enum.RoleSeeker))
	if err != nil {
		return "", err
	}

	return token, nil
}
