// Custom http. Wrap some http usage in this class.
// MAYBE TODO: Wrap [net/http] completely.
package chttp

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile(
	"^/($|viewer/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$|main/$|login/$|login_process/$|register/$|register_process/$|admin/$)",
)

// Validate each handler.
func validateHandler(
	fn func(http.ResponseWriter, *http.Request),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.Error(w, "INVALID PATH", http.StatusInternalServerError)
		} else {
			fn(w, r)
		}
	}
}
