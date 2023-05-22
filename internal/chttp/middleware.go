package chttp

import (
	"log"
	"mag/internal/chttp/logger"
	"net/http"
)

func NewLogger(l log.Logger) func(http.Handler) http.Handler {
	return logger.NewLogger(l)
}

/*
   func NewAuthenticator(u mag.User) func(http.Handler) http.Handler {
   return auth.NewAuthenticator(u)
   }
*/
