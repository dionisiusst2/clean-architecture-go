package interactor

import (
	"net/http"
)

func (intr *authInteractor) Logout(w http.ResponseWriter) {
	intr.cookieHandler.Delete(w, "token")
}
