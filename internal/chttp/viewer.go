package chttp

import (
	"mag"
	"mag/service"
	"net/http"

	"github.com/google/uuid"
)

func ViewHandler(
	s service.Service,
	renderFun func(http.ResponseWriter, *mag.Magazine) error,
	emptyFun func(http.ResponseWriter, error) error,
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
			emptyFun(w, err)
			return
		}
		renderFun(w, m)
	}
	return validateHandler(fn)
}
