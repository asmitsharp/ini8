// internal/repository/registration.go
package repository

import (
	"context"
	"ini8/internal/model"

	"gorm.io/gorm"
)

// RegistrationRepository defines the interface for registration database operations
type RegistrationRepository interface {
	Create(ctx context.Context, reg *model.Registration) error
	Get(ctx context.Context, id uint) (*model.Registration, error)
	List(ctx context.Context) ([]model.Registration, error)
	Update(ctx context.Context, reg *model.Registration) error
	Delete(ctx context.Context, id uint) error
}

// registrationRepository implements the RegistrationRepository interface
type registrationRepository struct {
	db *gorm.DB
}

// NewRegistrationRepository creates a new instance of registrationRepository
func NewRegistrationRepository(db *gorm.DB) RegistrationRepository {
	return &registrationRepository{db: db}
}

// Create implements the Create method of RegistrationRepository
func (r *registrationRepository) Create(ctx context.Context, reg *model.Registration) error {
	return r.db.WithContext(ctx).Create(reg).Error
}

// Get implements the Get method of RegistrationRepository
func (r *registrationRepository) Get(ctx context.Context, id uint) (*model.Registration, error) {
	var reg model.Registration
	if err := r.db.WithContext(ctx).First(&reg, id).Error; err != nil {
		return nil, err
	}
	return &reg, nil
}

// List implements the List method of RegistrationRepository
func (r *registrationRepository) List(ctx context.Context) ([]model.Registration, error) {
	var regs []model.Registration
	if err := r.db.WithContext(ctx).Find(&regs).Error; err != nil {
		return nil, err
	}
	return regs, nil
}

// Update implements the Update method of RegistrationRepository
func (r *registrationRepository) Update(ctx context.Context, reg *model.Registration) error {
	return r.db.WithContext(ctx).Save(reg).Error
}

// Delete implements the Delete method of RegistrationRepository
func (r *registrationRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Registration{}, id).Error
}
