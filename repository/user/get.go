package repository

import (
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/google/uuid"
)

func (r *userRepository) GetAllUser() ([]*domain.User, errors.HttpError) {
	var users []*domain.User
	err := r.db.Find(&users).Error

	if err != nil {
		return nil, errors.NewHttpError("user_repository.FindAll", http.StatusInternalServerError, err.Error())
	}

	return users, nil
}

func (r *userRepository) GetByID(uuid uuid.UUID) (*domain.User, errors.HttpError) {
	var user domain.User

	err := r.db.First(&user, uuid).Error
	if err != nil {
		return &user, errors.NewHttpError("user_repository.FindByID", http.StatusInternalServerError, err.Error())
	}

	return &user, nil
}
