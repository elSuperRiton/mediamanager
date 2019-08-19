package middlewares

import (
	"net/http"
	"time"
)

const responseTimeHeader = "X-Mediamanager-Time"

// requestTimerInterceptor implements the http.ResponseWriter
// interface and keeps a reference to the original http.ResponseWriter
// in order to write the request time before request compliton
type requestTimerInterceptor struct {
	startTime time.Time
	w         http.ResponseWriter
}

func newRequestTimer(w http.ResponseWriter) *requestTimerInterceptor {
	return &requestTimerInterceptor{
		startTime: time.Now(),
		w:         w,
	}
}

func (r *requestTimerInterceptor) Write(b []byte) (int, error) {
	return r.w.Write(b)
}

func (r *requestTimerInterceptor) Header() http.Header {
	return r.w.Header()
}

func (r *requestTimerInterceptor) WriteHeader(code int) {
	duration := time.Now().Sub(r.startTime)
	r.w.Header().Set(responseTimeHeader, duration.String())
	r.w.WriteHeader(code)
}

// RequestTimer intercepts request and writes out the time
// it took for them to complete in the header
func RequestTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqTimer := newRequestTimer(w)
		next.ServeHTTP(reqTimer, r)
	})
}
