package repository

import (
	"fmt"
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
)

func (r *authRepository) GetUserByEmail(email string) (domain.User, errors.HttpError) {
	var user domain.User
	r.db.Where("email = ?", email).Find(&user)
	if (user == domain.User{}) {
		return user, errors.NewHttpError("auth_repository.GetUserByEmail", http.StatusNotFound, fmt.Sprintf("user with email %v not found", email))
	}

	return user, nil
}
