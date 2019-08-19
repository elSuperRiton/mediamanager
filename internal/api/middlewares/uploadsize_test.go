package middlewares

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elSuperRiton/mediamanager/internal/pkg/config"
)

func TestMaxUploadSize(t *testing.T) {

	var maxAllowedSize int64 = 1 * 1024 * 1024 // 1Mo

	tests := []struct {
		imageName      string
		maxAllowedSize int64
		expected       int
	}{
		{
			imageName:      "big_img.jpg",
			maxAllowedSize: maxAllowedSize, // 1Mo
			expected:       http.StatusRequestEntityTooLarge,
		},
		{
			imageName:      "small_img.jpg",
			maxAllowedSize: maxAllowedSize, // 1Mo
			expected:       http.StatusOK,
		},
	}

	for _, test := range tests {

		config.Conf.MaxUploadSize = test.maxAllowedSize
		file, err := ioutil.ReadFile("./testdata/" + test.imageName)
		if err != nil {
			t.Errorf("error reading test file %v : %v", test.imageName, err)
		}

		// Prepare a form that you will submit to that URL.
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fileWriter, _ := w.CreateFormFile("file", test.imageName)
		fileWriter.Write(file)

		// Create test request on media mux
		req, err := http.NewRequest("POST", "/medias", &b)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Content-Type", w.FormDataContentType())

		// Create a ResponseRecorder to check against result
		rr := httptest.NewRecorder()

		// Create a stub hanlder calling WriteHeader method
		handler := NewMaxUploadSize(&MaxUploadSizeOptions{
			AllowedFileType: []string{"jpg"},
			Size:            test.maxAllowedSize,
		})

		h := handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		h.ServeHTTP(rr, req)

		if code := rr.Code; code != test.expected {
			t.Errorf(
				"expected response code of file upload %v to be %v, got %v",
				test.imageName,
				test.expected,
				code,
			)
		}
	}
}
