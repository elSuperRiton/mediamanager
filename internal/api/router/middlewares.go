package router

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/elSuperRiton/mediamanager/internal/api/middlewares"
	"github.com/go-chi/cors"
)

// getMaxUploadSizeMiddleware is a helper function returning a middleware for limiting
// file upload size
func getMaxUploadSizeMiddleware(maxSize int64, typesAllowed []string) middlewares.MiddlewareFunc {
	return middlewares.NewMaxUploadSize(&middlewares.MaxUploadSizeOptions{
		AllowedFileType: typesAllowed,
		Size:            maxSize,
	})
}

// getContentTypeMiddleware is a helper function returning a middleware for against
// Content-Type
func getContentTypeMiddleware(contentType string, strictCheck bool) middlewares.MiddlewareFunc {
	return middlewares.NewContentType(&middlewares.ContentTypeOptions{
		ContentType: contentType,
		Strict:      strictCheck,
	})
}

// getJWTMiddleware is a helper function returning a JWT middleware for authentication
func getJWTMiddleware() middlewares.MiddlewareFunc {
	return middlewares.NewJWT(&middlewares.JWTConfig{
		SigningKey: []byte("test"),
		HeaderKey:  []byte("Authorization"),
	})
}

// setCors is a helper function returning chi Cors middleware
func getCorsMiddleware() func(next http.Handler) http.Handler {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	return crs.Handler
}

// getAssetHandler is a helper function returning a fileHandler func
func getAssetHandler() (fileHandler http.HandlerFunc) {
	// Setup our assetHandler and point it to our static build location
	fs := http.StripPrefix("/", http.FileServer(rice.MustFindBox("../../../public/web").HTTPBox()))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
