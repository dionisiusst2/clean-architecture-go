package interactor

import "github.com/dionisiusst2/bakery-id/utils/errors"

func (intr *userInteractor) DeleteByID(id string) errors.HttpError {
	uuid, err := intr.parseToUUID(id)
	if err != nil {
		return err.WithOperation("user_interactor.DeleteByID")
	}

	err = intr.userRepo.DeleteByID(uuid)
	if err != nil {
		return err.WithOperation("user_interactor.DeleteByID")
	}

	return nil
}
