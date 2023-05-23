package chttp

import (
	"mag/service"
	"net/http"
)

func AdminHandler(
	s service.Service,
	renderFun func(http.ResponseWriter, bool) error,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		renderFun(w, true) // Middleware does auth.
	}
	return validateHandler(fn)
}
