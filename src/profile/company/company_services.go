package company

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
)

type CompanyService interface {
	Signup(dto dto.SignupDto) (string, error)
	Signin(dto dto.SigninDto) (string, error)
}

type Service struct {
	repo CompanyRepo
}

func NewService(repo CompanyRepo) CompanyService {
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

	company := model.CompanyProfile{
		ID:          uuid.New(),
		CompanyName: dto.Name,
		Email:       dto.Email,
		Password:    hashedPassword,
	}

	if err := s.repo.Create(company); err != nil {
		return "", err
	}

	token, err := helper.GenerateJwt(company.ID.String(), company.Email, string(enum.RoleCompany))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Signin(dto dto.SigninDto) (string, error) {
	company, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckPasswordHash(dto.Password, company.Password) {
		return "", fmt.Errorf("incorrect credentials")
	}

	token, err := helper.GenerateJwt(company.ID.String(), company.Email, string(enum.RoleCompany))
	if err != nil {
		return "", err
	}

	return token, nil
}
