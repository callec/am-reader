package logger

import (
	"log"
	"net/http"
)

// Logger middleware factory.
func NewLogger(l log.Logger) func(http.Handler) http.Handler {
	return basicLogger(l)
}
