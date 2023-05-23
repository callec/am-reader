package chttp

import (
	"mag/service"
	"net/http"
)

func LoginProcessHandler(
	s service.Service,
) http.HandlerFunc {
	a := newAuthoriser(s)
	return validateHandler(a.LoginHandler)
}

func LoginHandler(
	s service.Service,
	renderFun func(http.ResponseWriter) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := renderFun(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return validateHandler(fn)
}

func RegisterProcessHandler(
	s service.Service,
) http.HandlerFunc {
	a := newAuthoriser(s)
	return validateHandler(a.RegisterHandler)
}

func RegisterHandler(
	s service.Service,
	renderFun func(http.ResponseWriter) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := renderFun(w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return validateHandler(fn)
}
