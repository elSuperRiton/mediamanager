package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/elSuperRiton/mediamanager/pkg/models"
)

type (
	SetMediaContextOptions struct {
		Conf *models.MediaManagerConfig
	}
	setMediaContext struct {
		conf *models.MediaManagerConfig
	}
)

func NewSetMediaContext(options *SetMediaContextOptions) MiddlewareFunc {
	mdl := &setMediaContext{
		conf: options.Conf,
	}
	return mdl.Handler
}

func (s *setMediaContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fullPath := r.URL.Path
		splitPath := strings.Split(fullPath, "/")

		reqWithContext := r.WithContext(context.WithValue(r.Context(), "uploader", splitPath[len(splitPath)-1]))
		next.ServeHTTP(w, reqWithContext)
	})
}
