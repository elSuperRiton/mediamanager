package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthentify(t *testing.T) {
	testSuite := []struct {
		body               string
		expectedStatusCode int
	}{
		{
			body: `
				{
					"email": "test@test.com",
					"password": "pwd"
				}
			`,
			expectedStatusCode: http.StatusOK,
		},
		{
			body: `
				{
					"email": "bademail@bademail.com",
					"password": "badpassword"
				}
			`,
			expectedStatusCode: http.StatusUnauthorized,
		},
		{
			body: `
				{
					"email": "wrongjsonformat@wrongjsonformat.com",
					"password": "wrongjsonformat
				}
			`,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testSuite {
		req, err := http.NewRequest("POST", "/medias", strings.NewReader(test.body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Authentify)
		handler.ServeHTTP(rr, req)

		if statusCode := rr.Result().StatusCode; statusCode != test.expectedStatusCode {
			t.Errorf("expected status code to be %v, got %v", test.expectedStatusCode, statusCode)
		}
	}
}
