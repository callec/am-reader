package chttp

import (
	"mag"
	"mag/internal/chttp/auth"
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
	d mag.Service,
) Authoriser {
	if authoriser == nil {
		authoriser = auth.NewBasicAuthoriser(d)
	}
	return authoriser
}
