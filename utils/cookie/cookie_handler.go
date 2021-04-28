package cookie

import (
	"net/http"
	"os"
	"time"

	"github.com/dionisiusst2/clean-architecture-go/utils/errors"
	sc "github.com/gorilla/securecookie"
)

type Handler interface {
	Get(*http.Request, string) (string, errors.HttpError)
	Add(http.ResponseWriter, string, string) errors.HttpError
	Delete(http.ResponseWriter, string)
}

type cookieHandler struct {
	secureCookie *sc.SecureCookie
}

func NewHandler() Handler {
	hashKey := []byte(os.Getenv("COOKIE_HASH"))
	blockKey := []byte(os.Getenv("COOKIE_BLOCK"))
	secureCookie := sc.New(hashKey, blockKey)

	return &cookieHandler{secureCookie}
}

func (c *cookieHandler) Get(r *http.Request, name string) (string, errors.HttpError) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", errors.NewHttpError("cookie_handler.GetCookie", http.StatusInternalServerError, err.Error())
	}

	var value string
	err = c.secureCookie.Decode(name, cookie.Value, &value)
	if err != nil {
		return "", errors.NewHttpError("cookie_handler.GetCookie", http.StatusInternalServerError, err.Error())
	}

	return value, nil
}

func (c *cookieHandler) Add(w http.ResponseWriter, name string, value string) errors.HttpError {
	encoded, err := c.secureCookie.Encode(name, value)
	if err != nil {
		return errors.NewHttpError("cookie_handler.AddCookie", http.StatusInternalServerError, err.Error())
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    encoded,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	return nil
}

func (c *cookieHandler) Delete(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Millisecond),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}
