package repository

import (
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
)

func (r *authRepository) Register(userData domain.User) errors.HttpError {
	err := r.db.Create(&userData).Error
	if err != nil {
		return errors.NewHttpError("auth_repository.Register", http.StatusInternalServerError, err.Error())
	}

	return nil
}
