package controller

import (
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/gin-gonic/gin"
)

type AppController struct {
	User
	Auth
}

func handleError(c *gin.Context, err errors.HttpError) {
	err.PrintFailedOperations()
	c.JSON(err.GetStatusCode(), gin.H{
		"error": err,
	})
}
