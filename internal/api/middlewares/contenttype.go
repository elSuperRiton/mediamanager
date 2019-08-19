package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/elSuperRiton/mediamanager/pkg/utils"
)

type (
	// ContentTypeOptions allows for intanciating the ContentType
	// middleware with specific configuration
	// It is intended to be passed through the NewContentType function
	ContentTypeOptions struct {
		ContentType string
		Strict      bool
	}
	// contentType defines the instanciable part of the
	// contentType middleware
	contentType struct {
		contentType string
		strict      bool
	}
)

// NewContentType allocates a new Contentype and returns the handler
// associated with it
func NewContentType(options *ContentTypeOptions) MiddlewareFunc {
	ct := &contentType{
		contentType: options.ContentType,
		strict:      options.Strict,
	}

	return ct.Hanlder
}

// isOfContentType simply verifies that the request Content-Type header
// matched the one provided in the configuration
// If strict set to true then a total equality match will be performed else
// only the first part of the header will be checked against configuration
func (m *contentType) isOfContentType(r *http.Request) bool {
	contentTypeHeader := r.Header.Get("Content-Type")

	if !m.strict {
		return strings.Split(contentTypeHeader, ";")[0] == m.contentType
	}

	return contentTypeHeader == m.contentType
}

// Hanlder is a middleware enforcing a specific Content-Type header
// It simply checks that the request passed in has a specific header set and
// returns a http.StatusBadRequest in case it isn't
func (m *contentType) Hanlder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ok := m.isOfContentType(r); !ok {
			utils.RenderErr(
				w,
				r,
				fmt.Sprintf("Request should be of type %v, got %v", m.contentType, r.Header.Get("Content-Type")),
				http.StatusBadRequest,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
