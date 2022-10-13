package middlewere

import (
	"belajar-restful-api/helper"
	"belajar-restful-api/model/web"
	"net/http"
)

type AuthMiddlewere struct {
	Handler http.Handler
}

func NewAuthMiddlewere(handler http.Handler) *AuthMiddlewere {
	return &AuthMiddlewere{Handler: handler}
}

func (middlewere *AuthMiddlewere) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") == "kifeb" {
		middlewere.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Api Key Salah",
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
