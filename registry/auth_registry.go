package registry

import (
	"github.com/dionisiusst2/clean-architecture-go/interface/controller"
	"github.com/dionisiusst2/clean-architecture-go/interface/middleware"
	repository "github.com/dionisiusst2/clean-architecture-go/repository/auth"
	interactor "github.com/dionisiusst2/clean-architecture-go/usecase/interactor/auth"
	cookie "github.com/dionisiusst2/clean-architecture-go/utils/cookie"
	"github.com/dionisiusst2/clean-architecture-go/utils/token"
)

func (r *registry) NewAuthController() controller.Auth {
	return controller.NewAuthController(r.NewAuthInteractor())
}

func (r *registry) NewAuthInteractor() interactor.Auth {
	return interactor.NewAuthInteractor(r.NewAuthRepository(), r.NewCookieHandler(), r.NewTokenHandler())
}

func (r *registry) NewAuthRepository() repository.Auth {
	return repository.NewAuthRepository(r.db)
}

func (r *registry) NewAuthMiddleware() middleware.Auth {
	return middleware.NewAuthMiddleware(r.NewTokenHandler(), r.NewCookieHandler(), r.NewUserInteractor())
}

func (r *registry) NewTokenHandler() token.Handler {
	return token.NewHandler()
}
func (r *registry) NewCookieHandler() cookie.Handler {
	return cookie.NewHandler()
}
