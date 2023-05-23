package chttp

import (
	"mag/service"
	"net/http"
)

func LoginProcessHandler(
	s service.Service,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	a := newAuthoriser(s, emptyFun)
	return validateHandler(a.LoginHandler)
}

func LoginHandler(
	s service.Service,
	renderFun func(http.ResponseWriter) error,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := renderFun(w)
		if err != nil {
			emptyFun(w, err)
			return
		}
	}

	return validateHandler(fn)
}

func RegisterProcessHandler(
	s service.Service,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	a := newAuthoriser(s, emptyFun)
	return validateHandler(a.RegisterHandler)
}

func RegisterHandler(
	s service.Service,
	renderFun func(http.ResponseWriter) error,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := renderFun(w)
		if err != nil {
			emptyFun(w, err)
			return
		}
	}

	return validateHandler(fn)
}
