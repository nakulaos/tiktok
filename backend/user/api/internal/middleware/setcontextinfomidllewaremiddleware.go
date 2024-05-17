package middleware

import (
	"net/http"
	"tiktok/common/middleware"
)

type SetContextInfoMidllewareMiddleware struct {
}

func NewSetContextInfoMidllewareMiddleware() *SetContextInfoMidllewareMiddleware {
	return &SetContextInfoMidllewareMiddleware{}
}

func (m *SetContextInfoMidllewareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	//return func(w http.ResponseWriter, r *http.Request) {
	//	// TODO generate middleware implement function, delete after code implementation
	//	resource := r.Header.Get("Accept-Language")
	//	if resource == "" {
	//		resource = "en"
	//	}
	//	r = r.WithContext(context.WithValue(r.Context(), "resource", resource))
	//	// Passthrough to next handler if need
	//
	//	next(w, r)
	//}

	return middleware.SetContextInfoHandle(next)
}
