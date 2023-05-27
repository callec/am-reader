package chttp

import (
	"mag"
	"net/http"
)

func LoginProcessHandler(
	s mag.Service,
) http.HandlerFunc {
	a := newAuthoriser(s)
	return validateHandler(a.LoginHandler)
}

func LoginHandler(
	s mag.Service,
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
	s mag.Service,
) http.HandlerFunc {
	a := newAuthoriser(s)
	return validateHandler(a.RegisterHandler)
}

func RegisterHandler(
	s mag.Service,
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
