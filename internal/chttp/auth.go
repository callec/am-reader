package chttp

import "net/http"

type Authoriser interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	// A middleware used to make sure that a user is logged for certain pages.
	RequireLogin(next http.Handler) http.Handler
}

func NewAuthoriser() *Authoriser {
	return nil
}
