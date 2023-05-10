package chttp

import (
	"mag"
	"mag/service"
	"net/http"
)

func HomeHandler(
	s service.Service,
	renderFun func(http.ResponseWriter, []*mag.Magazine) error,
	emptyFun func(http.ResponseWriter, error) error,
) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		offset := 0
		limit := 20

		ctx, cf := getTimedContext(untilTimeOut)
		defer cf()
		ms, err := s.ListMagazines(ctx, limit, offset)
		if err != nil {
			emptyFun(w, err)
			return
		}
		renderFun(w, ms)
	}
	return validateHandler(fn)
}
