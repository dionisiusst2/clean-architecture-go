package token

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
)

type Handler interface {
	NewClaims(userID string) jwt.MapClaims
	GenerateSignedTokenWithClaims(jwt.MapClaims) (string, errors.HttpError)
	ExtractUserIDFromToken(string) (string, errors.HttpError)
}

type tokenHandler struct{}

func NewHandler() Handler {
	return &tokenHandler{}
}

func (t *tokenHandler) NewClaims(userID string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["ID"] = userID
	claims["exp"] = time.Now().Add(2 * time.Hour)

	return claims
}

func (t *tokenHandler) GenerateSignedTokenWithClaims(claims jwt.MapClaims) (string, errors.HttpError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.NewHttpError("token_handler.GenerateSignedTokenWithClaims", http.StatusInternalServerError, err.Error())
	}

	return signedToken, nil
}

func (t *tokenHandler) ExtractUserIDFromToken(tokenString string) (string, errors.HttpError) {
	var err error
	var token *jwt.Token
	if token, err = jwt.Parse(tokenString, keyFunc); err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims["ID"].(string), nil
		}
	}
	return "", errors.NewHttpError("token_handler.ExtractUserFromToken", http.StatusBadRequest, err.Error())
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("JWT_SECRET")), nil
}
