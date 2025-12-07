package jobseeker

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/ito-company/jobsito-service/helper"
	"github.com/ito-company/jobsito-service/src/dto"
	"github.com/ito-company/jobsito-service/src/enum"
	"github.com/ito-company/jobsito-service/src/model"
	"github.com/jinzhu/copier"
)

type JobSeekerService interface {
	Signup(dto dto.SignupDto) (string, error)
	Signin(dto dto.SigninDto) (string, error)
	InternSignin(dto dto.InternSigninDto) (string, error)
	Update(email string, input JobSeekerUpdateDto) (*dto.JobSeekerResponse, error)
	FindByEmail(email string) (*dto.JobSeekerResponse, error)
	FindById(id string) (*dto.JobSeekerResponse, error)
	SoftDelete(id string) error
	FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.JobSeekerResponse], error)
	AddTagToJobSeeker(jobSeekerId, tagId, proficiency string) error
	RemoveTagFromJobSeeker(jobSeekerId, tagId string) error
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

func (s *Service) InternSignin(dto dto.InternSigninDto) (string, error) {
	seeker, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckPasswordHash(dto.Password, seeker.Password) {
		return "", fmt.Errorf("incorrect credentials")
	}

	// Verificar si el pasante está asignado a la pasantía
	isInInternship, err := s.repo.IsJobSeekerInInternship(seeker.ID.String(), dto.InternshipId)
	if err != nil {
		return "", fmt.Errorf("error verifying internship: %w", err)
	}

	if !isInInternship {
		return "", fmt.Errorf("you are not assigned to this internship")
	}

	token, err := helper.GenerateJwt(seeker.ID.String(), seeker.Email, string(enum.RoleSeeker))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Update(email string, input JobSeekerUpdateDto) (*dto.JobSeekerResponse, error) {
	jobSeeker, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("jobseeker not found: %w", err)
	}

	opt := copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	}

	if err := copier.CopyWithOption(&jobSeeker, &input, opt); err != nil {
		return nil, fmt.Errorf("failed to copy data: %w", err)
	}

	if err := s.repo.Update(jobSeeker); err != nil {
		return nil, fmt.Errorf("failed to update jobseeker: %w", err)
	}

	dto := dto.JobSeekerToDto(&jobSeeker)
	return &dto, nil
}

func (s *Service) FindByEmail(email string) (*dto.JobSeekerResponse, error) {
	jobSeeker, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("jobseeker not found: %w", err)
	}
	dto := dto.JobSeekerToDto(&jobSeeker)
	return &dto, nil
}

func (s *Service) FindById(id string) (*dto.JobSeekerResponse, error) {
	jobSeeker, err := s.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("jobseeker not found: %w", err)
	}
	dto := dto.JobSeekerToDto(&jobSeeker)
	return &dto, nil
}

func (s *Service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *Service) FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[dto.JobSeekerResponse], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := dto.JobSeekerToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[dto.JobSeekerResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) AddTagToJobSeeker(jobSeekerId, tagId, proficiency string) error {
	return s.repo.AddTagToJobSeeker(jobSeekerId, tagId, proficiency)
}

func (s *Service) RemoveTagFromJobSeeker(jobSeekerId, tagId string) error {
	return s.repo.RemoveTagFromJobSeeker(jobSeekerId, tagId)
}
