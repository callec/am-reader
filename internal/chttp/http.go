// Custom http. Wrap all http usage in this class.
// TODO(kabe): Implement. Perhaps wrap net/http completely?
package chttp

import (
	"mag/magazine"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile(
	"^/($|viewer/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$|main/$)",
)

func MakeHandler(
	fn func(http.ResponseWriter, *http.Request, magazine.Service),
	s magazine.Service,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Error(w, "INVALID PATH", http.StatusInternalServerError)
		} else {
			fn(w, r, s)
		}
	}
}
