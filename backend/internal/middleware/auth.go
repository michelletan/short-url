package middleware

import (
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check JWT/cookie

		next.ServeHTTP(w, r)
	})
}
