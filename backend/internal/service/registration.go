package service

import (
	"context"
	"ini8/internal/model"
	"ini8/internal/repository"
	"ini8/pkg/validator"
)

type RegistrationService interface {
	Create(ctx context.Context, reg *model.Registration) error
	Get(ctx context.Context, id uint) (*model.Registration, error)
	List(ctx context.Context) ([]model.Registration, error)
	Update(ctx context.Context, reg *model.Registration) error
	Delete(ctx context.Context, id uint) error
}

type registrationService struct {
	repo repository.RegistrationRepository
}

func NewRegistrationService(repo repository.RegistrationRepository) RegistrationService {
	return &registrationService{repo: repo}
}

func (s *registrationService) Create(ctx context.Context, reg *model.Registration) error {
	if err := validator.ValidateRegistration(reg); err != nil {
		return err
	}
	return s.repo.Create(ctx, reg)
}

func (s *registrationService) Get(ctx context.Context, id uint) (*model.Registration, error) {
	return s.repo.Get(ctx, id)
}

func (s *registrationService) List(ctx context.Context) ([]model.Registration, error) {
	return s.repo.List(ctx)
}

func (s *registrationService) Update(ctx context.Context, reg *model.Registration) error {
	if err := validator.ValidateRegistration(reg); err != nil {
		return err
	}
	return s.repo.Update(ctx, reg)
}

func (s *registrationService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
