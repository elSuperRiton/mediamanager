package models

import (
	"mime/multipart"
	"plugin"
)

const (
	// UploaderTypeS3 is the type that needs to be associated with s3 uploader
	UploaderTypeS3 = "s3"
	// UploaderTypeFileSystem is the type that needs to be associated with fs uploader
	UploaderTypeFileSystem = "fs"
	// UploaderTypeCloudStorage is the type that needs to be associated with cloud storage
	UploaderTypeCloudStorage = "cloudStorage"
)

// UploaderCommons defines common fields between uploaders
// please make sure to use the `mapstructure:",squash"` when
// embedding UploaderCommons within new uploaders
type UploaderCommons struct {
	Type        string                                                           `yaml:"type" validate:"required"`
	Path        string                                                           `yaml:"path" validate:"required"`
	PluginName  string                                                           `yaml:"plugin"`
	Plugin      plugin.Plugin                                                    `yaml:"-"`
	GetFileName func(file *multipart.FileHeader) (newFileName string, err error) `yaml:"-"`
}
