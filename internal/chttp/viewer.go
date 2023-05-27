package chttp

import (
	"mag"
	"net/http"

	"github.com/google/uuid"
)

func ViewHandler(
	s mag.Service,
	renderFun func(http.ResponseWriter, *mag.Magazine) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.URL.Path[len("/viewer/"):])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx, cf := getTimedContext(untilTimeOut)
		defer cf()
		m, err := s.GetMagazine(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderFun(w, m)
	}
	return validateHandler(fn)
}
