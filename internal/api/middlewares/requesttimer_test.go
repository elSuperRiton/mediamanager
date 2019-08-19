package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// GetTestHandler returns a http.HandlerFunc for testing http middleware
func getTestHandler() http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
}

func TestRequestTimer(t *testing.T) {
	t.Run("testing request timer writes to header when calling WriteHeader", func(t *testing.T) {
		// Create test request on root mux
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a ResponseRecorder to check against result
		rr := httptest.NewRecorder()

		// Create a stub hanlder calling WriteHeader method
		handler := RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		// Serve and record request
		handler.ServeHTTP(rr, req)

		if timerHeader := rr.Header().Get(responseTimeHeader); timerHeader == "" {
			t.Errorf("request should have %v attached to Headers", responseTimeHeader)
		}
	})

	t.Run("testing no request timer header is set if no call to WriteHeader occures", func(t *testing.T) {
		// Create test request on root mux
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a ResponseRecorder to check against result
		rr := httptest.NewRecorder()

		// Create a stub hanlder calling WriteHeader method
		handler := RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

		// Serve and record request
		handler.ServeHTTP(rr, req)

		if timerHeader := rr.Header().Get(responseTimeHeader); timerHeader != "" {
			t.Errorf("request should have %v attached to Headers", responseTimeHeader)
		}
	})
}
