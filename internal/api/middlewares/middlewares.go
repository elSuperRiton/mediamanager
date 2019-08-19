package middlewares

import "net/http"

// MiddlewareFunc is a helper type for defining middleware
// functions
type MiddlewareFunc func(next http.Handler) http.Handler
