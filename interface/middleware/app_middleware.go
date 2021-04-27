package middleware

import (
	"github.com/dionisiusst2/bakery-id/utils/errors"
	"github.com/gin-gonic/gin"
)

type AppMiddleware struct {
	Auth
}

func handleError(c *gin.Context, err errors.HttpError) {
	err.PrintFailedOperations()
	c.JSON(err.GetStatusCode(), gin.H{
		"error": err,
	})
	c.Abort()
}
