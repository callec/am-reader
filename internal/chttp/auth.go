package chttp

import (
	"mag/internal/chttp/auth"
	"mag/service"
	"net/http"
)

type Authoriser interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	// A middleware used to make sure that a user is logged for certain pages.
	RequireLogin(next http.Handler) http.Handler
}

func NewAuthoriser(
	d service.Service,
	ef func(http.ResponseWriter, error) error,
) Authoriser {
	return auth.NewBasicAuthoriser(d, ef)
}
