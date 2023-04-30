package chttp

import (
	"mag/magazine"
	"net/http"

	"github.com/google/uuid"
)

func ViewHandler(
	w http.ResponseWriter,
	r *http.Request,
	s magazine.Service,
	renderFun func(http.ResponseWriter, *magazine.Magazine) error,
	emptyFun func(http.ResponseWriter, error) error,
) {
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
