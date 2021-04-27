package router

import (
	"github.com/dionisiusst2/bakery-id/interface/controller"
	"github.com/dionisiusst2/bakery-id/interface/middleware"
	"github.com/gin-gonic/gin"
)

func AddAuthRoutes(r *gin.RouterGroup, ctrl controller.Auth, m middleware.Auth) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", ctrl.Register)
		auth.POST("/login", ctrl.Login)
		auth.GET("/logout", m.Auth(), ctrl.Logout)
		auth.POST("/me", m.Auth(), ctrl.GetMe)
	}
}
