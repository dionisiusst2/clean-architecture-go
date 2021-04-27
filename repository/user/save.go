package repository

import (
	"net/http"

	"github.com/dionisiusst2/bakery-id/domain"
	"github.com/dionisiusst2/bakery-id/utils/errors"
)

func (r *userRepository) Save(user *domain.User) errors.HttpError {
	err := r.db.Save(user).Error
	if err != nil {
		return errors.NewHttpError("user_repository.UpdateByID", http.StatusInternalServerError, err.Error())
	}

	return nil
}
