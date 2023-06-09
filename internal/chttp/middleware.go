package chttp

import (
	"log"
	"mag"
	"mag/internal/chttp/errormw"
	"mag/internal/chttp/logger"
	"net/http"
)

func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return logger.NewLogger(l)
}

func NewAuth(
	d mag.Service,
) func(http.Handler) http.Handler {
	a := newAuthoriser(d)
	return a.RequireLogin
}

func NewError(ef func(http.ResponseWriter, error) error) func(http.Handler) http.Handler {
	return errormw.ErrorMiddleware(ef)
}
