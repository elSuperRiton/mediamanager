package handlers

import (
	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/uploader"
)

type Repository struct {
	conf      *models.MediaManagerConfig
	uploaders map[string]uploader.Uploader
}

var repository *Repository

// NewRepository needs to be called before using any handlers
// It sets up the repository with needed
func NewRepository(conf *models.MediaManagerConfig) {
	repository = &Repository{
		conf: conf,
	}
}
