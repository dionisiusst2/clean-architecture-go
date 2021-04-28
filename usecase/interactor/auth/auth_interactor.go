package interactor

import (
	"net/http"

	"github.com/dionisiusst2/clean-architecture-go/domain"
	repository "github.com/dionisiusst2/clean-architecture-go/repository/auth"
	cookie "github.com/dionisiusst2/clean-architecture-go/utils/cookie"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	"github.com/dionisiusst2/clean-architecture-go/utils/token"
)

type Auth interface {
	Register(w http.ResponseWriter, user domain.User) (token string, err errors.HttpError)
	Login(w http.ResponseWriter, email string, password string) (token string, err errors.HttpError)
	Logout(w http.ResponseWriter)
}

type authInteractor struct {
	authRepo      repository.Auth
	cookieHandler cookie.Handler
	tokenHandler  token.Handler
}

func NewAuthInteractor(r repository.Auth, c cookie.Handler, t token.Handler) Auth {
	return &authInteractor{r, c, t}
}

func (intr *authInteractor) generateTokenAndSaveToCookie(w http.ResponseWriter, user domain.User) (string, errors.HttpError) {
	claims := intr.tokenHandler.NewClaims(user.ID.String())
	token, err := intr.tokenHandler.GenerateSignedTokenWithClaims(claims)
	if err != nil {
		return "", err.WithOperation("auth_interactor.generateTokenAndSaveToCookie")
	}

	err = intr.cookieHandler.Add(w, "token", token)
	if err != nil {
		return "", err.WithOperation("auth_interactor.generateTokenAndSaveToCookie")
	}

	return token, nil
}
