package interactor

import (
	"log"
	"net/http"
	"strings"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func (intr *userInteractor) UpdateByID(id string, userData domain.User) (domain.User, errors.HttpError) {
	err := intr.validateField(userData)
	if err != nil {
		return domain.User{}, err.WithOperation("user_interactor.UpdateByID")
	}

	uuid, err := intr.parseToUUID(id)
	if err != nil {
		return domain.User{}, err.WithOperation("user_interactor.UpdateByID")
	}

	user, err := intr.userRepo.GetByID(uuid)
	if err != nil {
		return domain.User{}, err.WithOperation("user_interactor.UpdateByID")
	}

	intr.updateField(user, userData)

	err = intr.userRepo.Save(user)
	if err != nil {
		return domain.User{}, err.WithOperation("user_interactor.UpdateByID")
	}

	return *user, nil
}

func (intr *userInteractor) validateField(user domain.User) errors.HttpError {
	err := user.Validate()
	if len(err) > 0 {
		return errors.NewHttpError("auth_interactor.validateField", http.StatusBadRequest, strings.Join(err[:], ", "))
	}

	return nil
}

func (intr *userInteractor) updateField(user *domain.User, newUserData domain.User) {
	user.Name = newUserData.Name
	user.Email = newUserData.Email
	user.Password = intr.getHashedPassword(newUserData.Password)
}

func (intr *userInteractor) getHashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
