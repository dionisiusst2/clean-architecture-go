package repository

import (
	"github.com/dionisiusst2/bakery-id/domain"
	"github.com/dionisiusst2/bakery-id/utils/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User interface {
	GetAllUser() ([]*domain.User, errors.HttpError)
	GetByID(uuid uuid.UUID) (*domain.User, errors.HttpError)
	Save(user *domain.User) errors.HttpError
	DeleteByID(uuid uuid.UUID) errors.HttpError
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &userRepository{db}
}
