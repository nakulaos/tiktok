package middleware

import (
	"context"
	"net/http"
)

func SetContextInfoHandle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		lang := r.Header.Get("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		r = r.WithContext(context.WithValue(r.Context(), "lang", lang))
		// Passthrough to next handler if need

		next(w, r)
	}
}
