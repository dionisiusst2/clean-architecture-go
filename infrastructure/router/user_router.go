package router

import (
	"github.com/dionisiusst2/bakery-id/interface/controller"
	"github.com/dionisiusst2/bakery-id/interface/middleware"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(r *gin.RouterGroup, ctrl controller.User, m middleware.Auth) {
	user := r.Group("/user")
	{
		user.GET("/", ctrl.FindAll)
		user.GET("/:id", m.Admin(), ctrl.FindByID)
		user.PUT("/:id", m.Auth(), ctrl.UpdateByID)
		user.DELETE("/:id", m.Admin(), ctrl.DeleteByID)
	}
}
