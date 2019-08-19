package middlewares

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/elSuperRiton/mediamanager/errors"
	"github.com/elSuperRiton/mediamanager/pkg/utils"
)

const exceededContentSize = "The file you're trying to upload is too large"

type (
	// MaxUploadSizeOptions defines the configuration
	// options available for the MaxUploadSize middleware
	MaxUploadSizeOptions struct {
		AllowedFileType []string
		Size            int64
	}

	// MaxUploadSize defines the instanciable part of the
	// MaxUploadSize middleware
	maxUploadSize struct {
		MaxUploadSizeOptions
		reader io.ReadCloser
		read   int64
	}
)

// NewMaxUploadSize returns a MaxUploadSizeMiddleware
//
// MaxUploadSizeMiddleware sets the maximum allowed size for a request body.
// In case the request body is above the limit it'll return a 413 error code
// The MaxUploadSizeMiddleware checks the body length based on the 'Content-Length'
// header and reading the body content
func NewMaxUploadSize(options *MaxUploadSizeOptions) MiddlewareFunc {
	maxUploadSize := &maxUploadSize{MaxUploadSizeOptions: *options}
	return maxUploadSize.Handler
}

// isContentLengthBelowLimit reads the Content-Length Header and checks
// if it's value is below the maximum allowed size
// If the Header value cannot be parsed to int64 it returns an error
func (m *maxUploadSize) isContentLengthBelowLimit(r *http.Request) (isvalid bool, parseErr error) {

	rawContentLengthVal := r.Header.Get("Content-Length")
	if rawContentLengthVal != "" {

		n, err := strconv.ParseInt(rawContentLengthVal, 10, 64)
		if err != nil {
			return false, err
		}

		if n < m.Size {
			return true, nil
		}
	}

	return true, nil
}

func (m *maxUploadSize) isBodyBelowLimit(r *http.Request, w http.ResponseWriter, pool *sync.Pool) (bodyBytes []byte, bodyReadingErr error) {
	// Limit size by reading actual body length
	// Here we call Reset before reading content in
	// order to make sure the instance of MaxUploadSize
	// is clean for usage
	mx := pool.Get().(*maxUploadSize)
	mx.Reset(r.Body)
	defer pool.Put(mx)

	b := make([]byte, m.Size)
	for {
		_, err := mx.Read(b)
		if err != nil {
			if err != io.EOF {
				mx.Close()
				return b, err
			}

			break
		}
	}

	return b, nil
}

// Handler is a middleware to enforce the maximum body size of a request
// It first checks the Content-Length header against the provided configuration
// prior to reading body into a buffer that cannot exceed the provided limit size
// It returns a http.StatusBadRequest in case the request body size is too large
func (m *maxUploadSize) Handler(next http.Handler) http.Handler {

	pool := maxUploadSizePool(m.MaxUploadSizeOptions)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// limit size based on content length
		// Note that as Content-Length can be tempered with we simply check
		// if it's value exceeds the max allowed size
		if ok, err := m.isContentLengthBelowLimit(r); !ok || err != nil {
			if err != nil {
				utils.RenderErr(w, r, err.Error(), http.StatusBadRequest)
				return
			}

			utils.RenderErr(w, r, exceededContentSize, http.StatusRequestEntityTooLarge)
			return
		}

		// Limit size by reading actual body length
		// Here we call Reset before reading content in
		// order to make sure the instance of MaxUploadSize
		// is clean for usage
		mx := pool.Get().(*maxUploadSize)
		mx.Reset(r.Body)
		defer pool.Put(mx)

		b, err := m.isBodyBelowLimit(r, w, &pool)
		if err != nil {
			switch err.(type) {
			case errors.ErrFileTooLarge:
				utils.RenderErr(
					w,
					r,
					err.Error(),
					http.StatusRequestEntityTooLarge,
				)
				return
			default:
				utils.RenderErr(
					w,
					r,
					err.Error(),
					http.StatusInternalServerError,
				)
				return
			}
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		next.ServeHTTP(w, r)
	})
}

// Read method reads from a slice of byte (here intended to be request body)
// and errors out if the number of bytes read exceeds the allowed body size
func (m *maxUploadSize) Read(b []byte) (n int, err error) {
	n, err = m.reader.Read(b)
	m.read += int64(n)

	if m.read > m.Size {
		return n, errors.NewErrFileTooLarge(m.Size)
	}

	return
}

// Close cleans up the memory allocated to reading MaxUploadSize reader
func (m *maxUploadSize) Close() error {
	return m.reader.Close()
}

// Reset allows for reseting MaxUploadSize before placing it back in the pool
func (m *maxUploadSize) Reset(readCloser io.ReadCloser) {
	m.reader = readCloser
	m.read = 0
}

// maxUploadSizePool return a new sync.Pool setup with an instance of
// MaxUploadSize middleware
// It enables us to save on memory by reusing instances of MaxUploadSize
func maxUploadSizePool(options MaxUploadSizeOptions) sync.Pool {
	return sync.Pool{
		New: func() interface{} {
			return &maxUploadSize{MaxUploadSizeOptions: options}
		},
	}
}
