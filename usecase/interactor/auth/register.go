package interactor

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (intr *authInteractor) Register(w http.ResponseWriter, user domain.User) (token string, err errors.HttpError) {
	err = intr.validateField(user)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Register")
	}

	_, err = intr.isEmailUnique(user.Email)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Register")
	}

	intr.populateWithIDAndHashedPassword(&user)

	err = intr.authRepo.Register(user)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Register")
	}

	token, err = intr.generateTokenAndSaveToCookie(w, user)
	if err != nil {
		return "", err.WithOperation("auth_interactor.Register")
	}

	return token, nil
}

func (intr *authInteractor) isEmailUnique(email string) (bool, errors.HttpError) {
	user, _ := intr.authRepo.GetUserByEmail(email)
	if (user != domain.User{}) {
		return false, errors.NewHttpError("auth_interactor.IsEmailUnique", http.StatusConflict, fmt.Sprintf("Email %v already exists", email))
	}

	return true, nil
}

func (intr *authInteractor) populateWithIDAndHashedPassword(user *domain.User) {
	user.ID = uuid.New()
	user.Password = intr.getHashedPassword(user.Password)
}

func (intr *authInteractor) getHashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func (intr *authInteractor) validateField(user domain.User) errors.HttpError {
	err := user.Validate()
	if len(err) > 0 {
		return errors.NewHttpError("auth_interactor.validateField", http.StatusBadRequest, strings.Join(err[:], ", "))
	}

	return nil
}
