package interactor

import (
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	repository "github.com/dionisiusst2/clean-architecture-go/repository/user"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/google/uuid"
)

type User interface {
	GetAllUser() ([]*domain.User, errors.HttpError)
	GetByID(id string) (*domain.User, errors.HttpError)
	UpdateByID(id string, userData domain.User) (domain.User, errors.HttpError)
	DeleteByID(id string) errors.HttpError
}

type userInteractor struct {
	userRepo repository.User
}

func NewUserInteractor(userRepo repository.User) User {
	return &userInteractor{userRepo}
}

func (intr *userInteractor) parseToUUID(ID string) (uuid.UUID, errors.HttpError) {
	uuid, err := uuid.Parse(ID)
	if err != nil {
		return uuid, errors.NewHttpError("user_interactor.parseToUUID", http.StatusBadRequest, err.Error())
	}

	return uuid, nil
}

func (intr *userInteractor) omitPasswordField(user *domain.User) {
	user.Password = ""
}
