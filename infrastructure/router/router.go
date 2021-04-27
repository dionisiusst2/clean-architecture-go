package router

import (
	"github.com/dionisiusst2/bakery-id/interface/controller"
	"github.com/dionisiusst2/bakery-id/interface/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(c controller.AppController, m middleware.AppMiddleware) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		AddUserRoutes(v1, c.User, m.Auth)
		AddAuthRoutes(v1, c.Auth, m.Auth)
	}

	return r
}
