package repository

import (
	"github.com/dionisiusst2/bakery-id/domain"
	"github.com/dionisiusst2/bakery-id/utils/errors"
	"gorm.io/gorm"
)

type Auth interface {
	Register(userData domain.User) errors.HttpError
	GetUserByEmail(email string) (domain.User, errors.HttpError)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) Auth {
	return &authRepository{db}
}
