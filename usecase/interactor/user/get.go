package interactor

import (
	"github.com/dionisiusst2/bakery-id/domain"
	"github.com/dionisiusst2/bakery-id/utils/errors"
)

func (intr *userInteractor) GetAllUser() ([]*domain.User, errors.HttpError) {
	users, err := intr.userRepo.GetAllUser()
	if err != nil {
		return nil, err.WithOperation("user_interactor.FindAll")
	}

	for _, user := range users {
		intr.omitPasswordField(user)
	}

	return users, nil
}

func (intr *userInteractor) GetByID(id string) (*domain.User, errors.HttpError) {

	uuid, err := intr.parseToUUID(id)
	if err != nil {
		return &domain.User{}, err.WithOperation("user_interactor.FindByID")
	}

	user, err := intr.userRepo.GetByID(uuid)
	if err != nil {
		return user, err.WithOperation("user_interactor.FindByID")
	}

	intr.omitPasswordField(user)

	return user, nil
}
