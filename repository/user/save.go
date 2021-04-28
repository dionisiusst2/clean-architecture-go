package repository

import (
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
)

func (r *userRepository) Save(user *domain.User) errors.HttpError {
	err := r.db.Save(user).Error
	if err != nil {
		return errors.NewHttpError("user_repository.UpdateByID", http.StatusInternalServerError, err.Error())
	}

	return nil
}
