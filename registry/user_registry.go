package registry

import (
	"github.com/dionisiusst2/clean-architecture-go/interface/controller"
	repository "github.com/dionisiusst2/clean-architecture-go/repository/user"
	interactor "github.com/dionisiusst2/clean-architecture-go/usecase/interactor/user"
)

func (r *registry) NewUserController() controller.User {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.User {
	return interactor.NewUserInteractor(r.NewUserRepository())
}

func (r *registry) NewUserRepository() repository.User {
	return repository.NewUserRepository(r.db)
}
