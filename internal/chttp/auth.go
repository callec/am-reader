package chttp

import (
	"mag/internal/chttp/auth"
	"mag/service"
	"net/http"
)

var authoriser Authoriser = nil

type Authoriser interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	// A middleware used to make sure that a user is logged for certain pages.
	RequireLogin(next http.Handler) http.Handler
}

func newAuthoriser(
	d service.Service,
) Authoriser {
	if authoriser == nil {
		authoriser = auth.NewBasicAuthoriser(d)
	}
	return authoriser
}
