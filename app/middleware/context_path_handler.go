package middleware

import (
	"net/http"
	"strings"
)

func ContextPathMiddleware(contextPath string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, contextPath) {
			http.NotFound(w, r)
			return
		}
		r.URL.Path = strings.TrimPrefix(r.URL.Path, contextPath)
		next.ServeHTTP(w, r) // Pass to the next handler
	})
}
