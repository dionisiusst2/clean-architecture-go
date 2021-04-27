package repository

import (
	"net/http"

	"github.com/dionisiusst2/bakery-id/domain"
	"github.com/dionisiusst2/bakery-id/utils/errors"
	"github.com/google/uuid"
)

func (r *userRepository) DeleteByID(uuid uuid.UUID) errors.HttpError {
	var user domain.User

	err := r.db.First(&user, uuid).Error
	if err != nil {
		return errors.NewHttpError("user_repository.DeleteByID", http.StatusInternalServerError, err.Error())
	}

	r.db.Delete(&domain.User{}, uuid)

	return nil
}
