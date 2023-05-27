package chttp

import (
	"mag"
	"net/http"
)

func HomeHandler(
	s mag.Service,
	renderFun func(http.ResponseWriter, []*mag.Magazine) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		offset := 0
		limit := 20

		ctx, cf := getTimedContext(untilTimeOut)
		defer cf()
		ms, err := s.ListMagazines(ctx, limit, offset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		renderFun(w, ms)
	}
	return validateHandler(fn)
}
