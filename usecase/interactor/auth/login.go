package interactor

import (
	"net/http"

	"github.com/dionisiusst2/bakery-id/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func (intr *authInteractor) Login(w http.ResponseWriter, email string, password string) (token string, err errors.HttpError) {
	user, err := intr.authRepo.GetUserByEmail(email)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Login")
	}

	passwordMatch := intr.isPasswordMatch(password, user.Password)
	if !passwordMatch {
		return "", errors.NewHttpError("auth_interactor.Login", http.StatusUnauthorized, "incorrect password.")
	}

	token, err = intr.generateTokenAndSaveToCookie(w, user)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Login")
	}

	return token, nil
}

func (intr *authInteractor) isPasswordMatch(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}

	return true
}
