package controller

import (
	"fmt"
	"net/http"

	"github.com/dionisiusst2/bakery-id/domain"
	interactor "github.com/dionisiusst2/bakery-id/usecase/interactor/user"
	"github.com/gin-gonic/gin"
)

type User interface {
	FindAll(*gin.Context)
	FindByID(*gin.Context)
	UpdateByID(*gin.Context)
	DeleteByID(*gin.Context)
}

type userController struct {
	intr interactor.User
}

func NewUserController(i interactor.User) User {
	return &userController{i}
}

func (ctrl *userController) FindAll(c *gin.Context) {
	users, err := ctrl.intr.GetAllUser()

	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func (ctrl *userController) FindByID(c *gin.Context) {
	id := c.Param("id")
	user, err := ctrl.intr.GetByID(id)

	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

func (ctrl *userController) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var userData domain.User
	c.ShouldBind(&userData)

	user, err := ctrl.intr.UpdateByID(id, userData)
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}

func (ctrl *userController) DeleteByID(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.intr.DeleteByID(id)
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("User with id %v deleted", id),
		})
	}
}
