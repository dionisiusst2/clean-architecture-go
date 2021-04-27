package controller

import (
	"net/http"

	"github.com/dionisiusst2/bakery-id/domain"
	interactor "github.com/dionisiusst2/bakery-id/usecase/interactor/auth"
	"github.com/dionisiusst2/bakery-id/utils/errors"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	Register(*gin.Context)
	Login(*gin.Context)
	Logout(*gin.Context)
	GetMe(*gin.Context)
}

type authController struct {
	intr interactor.Auth
}

func NewAuthController(i interactor.Auth) Auth {
	return &authController{i}
}

func (ctrl *authController) Register(c *gin.Context) {
	var userData domain.User
	c.ShouldBind(&userData)

	token, err := ctrl.intr.Register(c.Writer, userData)

	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func (ctrl *authController) Login(c *gin.Context) {
	var userData domain.User
	c.ShouldBind(&userData)

	token, err := ctrl.intr.Login(c.Writer, userData.Email, userData.Password)
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func (ctrl *authController) GetMe(c *gin.Context) {
	user, ok := c.Get("user")

	if !ok {
		err := errors.NewHttpError("auth_controller.GetMe", http.StatusInternalServerError, "internal server error.")
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": user.(domain.User),
		})
	}
}

func (ctrl *authController) Logout(c *gin.Context) {
	ctrl.intr.Logout(c.Writer)

	c.JSON(http.StatusOK, gin.H{
		"message": "cookie deleted",
	})
}
