package registry

import (
	"github.com/dionisiusst2/clean-architecture-go/interface/controller"
	"github.com/dionisiusst2/clean-architecture-go/interface/middleware"
	"gorm.io/gorm"
)

type Registry interface {
	NewAppController() controller.AppController
	NewAppMiddleware() middleware.AppMiddleware
}

type registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User: r.NewUserController(),
		Auth: r.NewAuthController(),
	}
}

func (r *registry) NewAppMiddleware() middleware.AppMiddleware {
	return middleware.AppMiddleware{
		Auth: r.NewAuthMiddleware(),
	}
}
