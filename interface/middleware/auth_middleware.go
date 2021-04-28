package middleware

import (
	"net/http"
	"strings"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	interactor "github.com/dionisiusst2/clean-architecture-go/usecase/interactor/user"
	cookie "github.com/dionisiusst2/clean-architecture-go/utils/cookie"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/dionisiusst2/clean-architecture-go/utils/token"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	Auth() gin.HandlerFunc
	Admin() gin.HandlerFunc
}

type authMiddleware struct {
	tokenHandler  token.Handler
	cookieHandler cookie.Handler
	userIntr      interactor.User
}

func NewAuthMiddleware(t token.Handler, c cookie.Handler, intr interactor.User) Auth {
	return &authMiddleware{t, c, intr}
}

func (m *authMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := m.getToken(c)
		if err != nil {
			handleError(c, err)
			return
		}

		userID, err := m.tokenHandler.ExtractUserIDFromToken(tokenString)
		if err != nil {
			handleError(c, err)
			return
		}

		user, err := m.userIntr.GetByID(userID)
		if err != nil {
			handleError(c, err)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (m *authMiddleware) Admin() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString, err := m.getToken(c)
		if err != nil {
			handleError(c, err)
			return
		}

		userID, err := m.tokenHandler.ExtractUserIDFromToken(tokenString)
		if err != nil {
			handleError(c, err)
			return
		}

		user, err := m.userIntr.GetByID(userID)
		if err != nil {
			handleError(c, err)
			return
		}

		if isAdmin, err := m.isRoleAdmin(user); !isAdmin || err != nil {
			handleError(c, err)
			return
		}

		c.Next()
	}
}

func (m *authMiddleware) getToken(c *gin.Context) (string, errors.HttpError) {
	var tokenString string

	if tokenString = m.getTokenFromAuthorizationHeader(c); tokenString == "" {
		tokenString = m.getTokenFromCookie(c)
	}

	if tokenString == "" {
		return "", errors.NewHttpError("auth_middleware.getToken", http.StatusUnauthorized, "unauthorized. please login first")
	}

	return tokenString, nil
}

func (m *authMiddleware) getTokenFromAuthorizationHeader(c *gin.Context) string {
	if header := c.Request.Header["Authorization"]; header != nil {
		return strings.Split(header[0], " ")[1]
	}

	return ""
}

func (m *authMiddleware) getTokenFromCookie(c *gin.Context) string {
	if tokenString, err := m.cookieHandler.Get(c.Request, "token"); err == nil {
		return tokenString
	}

	return ""
}

func (m *authMiddleware) isRoleAdmin(user *domain.User) (bool, errors.HttpError) {
	if user.Role != "admin" {
		return false, errors.NewHttpError("auth_middleware.isRoleAdmin", http.StatusForbidden, "you do not have access to this page.")
	}

	return true, nil
}
