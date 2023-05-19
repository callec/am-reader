package chttp

import (
	"log"
	"mag/internal/chttp/basiclogger"
	"net/http"
)

func NewLogger(l log.Logger) func(http.Handler) http.Handler {
	return basiclogger.BasicLogger(l)
}
