package chttp

import (
	"log"
	"mag/internal/chttp/logger"
	"mag/service"
	"net/http"
)

func NewLogger(l log.Logger) func(http.Handler) http.Handler {
	return logger.NewLogger(l)
}

func NewAuth(
	d service.Service,
	ef func(http.ResponseWriter, error) error,
) func(http.Handler) http.Handler {
	a := newAuthoriser(d, ef)
	return a.RequireLogin
}
