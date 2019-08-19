package middlewares_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elSuperRiton/mediamanager/internal/api/middlewares"
)

func TestNewContentType(t *testing.T) {
	testSuite := []struct {
		contentType        string
		reqContentType     string
		expectedStatusCode int
	}{
		{
			contentType:        "multipart/form-data",
			reqContentType:     "multipart/form-data",
			expectedStatusCode: http.StatusOK,
		},
		{
			contentType:        "application/json",
			reqContentType:     "multipart/form-data",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			contentType:        "application/x-www-form-urlencoded",
			reqContentType:     "application/x-www-form-urlencoded",
			expectedStatusCode: http.StatusOK,
		},
		{
			contentType:        "application/x-www-form-urlencoded",
			reqContentType:     "multipart/form-data",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			contentType:        "application/json",
			reqContentType:     "application/x-www-form-urlencoded",
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testSuite {

		req, err := http.NewRequest("POST", "/medias", bytes.NewReader([]byte{}))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", test.reqContentType)
		rr := httptest.NewRecorder()

		handler := middlewares.NewContentType(&middlewares.ContentTypeOptions{
			ContentType: test.contentType,
		})

		h := handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		h.ServeHTTP(rr, req)

		if status := rr.Result().StatusCode; status != test.expectedStatusCode {
			t.Errorf("wanted Status code to be %v, got %v", test.expectedStatusCode, status)
		}
	}
}
