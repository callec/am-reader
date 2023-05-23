package errormw

import (
	"fmt"
	"net/http"
)

func ErrorMiddleware(
	ef func(http.ResponseWriter, error) error,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					// Handle the panic and render an error response
					err := fmt.Errorf("Internal Server Error")
					if err := ef(w, err); err != nil {
						// Handle any error from the error renderer
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}

				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
