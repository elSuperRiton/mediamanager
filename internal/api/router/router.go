package router

import (
	"github.com/elSuperRiton/mediamanager/internal/api/handlers"
	"github.com/elSuperRiton/mediamanager/internal/api/middlewares"
	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/go-chi/chi"
)

const filesLocation = "../../../public/web"

var (
	conf *models.MediaManagerConfig
)

// New returns the media manager main handler built on top of
// chi router
func New(c *models.MediaManagerConfig) *chi.Mux {

	mainRouter := chi.NewMux()
	conf = c
	handlers.NewRepository(conf)

	// Global Mdls
	mainRouter.Use(middlewares.RequestTimer)
	mainRouter.Use(getCorsMiddleware())

	// Index
	mainRouter.Get("/", getAssetHandler())
	mainRouter.Get("/static/*", getAssetHandler())

	// Authentication routes
	mainRouter.Post("/auth", handlers.Authentify)

	// Authentified routes
	// All routes declared withing the authRouter are required to be authenticated
	mainRouter.Group(func(authRouter chi.Router) {
		// authRouter.Use(getJWTMiddleware())

		// Mount secondary routers
		authRouter.Mount("/medias", getMediasRouter())
	})

	return mainRouter
}

// getMediasRouter is a helper function returning routes under /medias
func getMediasRouter() (mediaRouter chi.Router) {
	mediaRouter = chi.NewMux()
	mediaRouter.Get("/", handlers.MediasFindAll)
	mediaRouter.Get("/url", handlers.MediasUploadURL)
	mediaRouter.Group(func(r chi.Router) {
		r.Use(getContentTypeMiddleware("multipart/form-data", false))
		// r.Use(getMaxUploadSizeMiddleware(1*1024*1024, []string{"jpg"}))
		r.Use(middlewares.NewSetMediaContext(&middlewares.SetMediaContextOptions{Conf: conf}))
		for _, uploader := range conf.Uploaders {
			path, _ := uploader["path"].(string)
			r.Post("/"+path, handlers.MediasUpload)
		}
	})

	return
}
