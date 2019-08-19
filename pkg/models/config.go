package models

import (
	"time"

	"github.com/elSuperRiton/mediamanager/pkg/uploader"
)

// MediaManagerConfig holds the configuration for the media manager api
type MediaManagerConfig struct {
	Port                 string                       `yaml:"port" validate:"required" env:"MEDIA_MANAGER_PORT" envDefault:":8080"`
	TimeOut              time.Duration                `yaml:"timeout" validate:"required" env:"MEDIA_MANAGER_TIMEOUT" envDefault:"30s"`
	MaxUploadSize        int64                        `yaml:"maxuploadsize" validate:"required" env:"MEDIA_MANAGER_MAX_UPLOAD" envDefault:"1024000"`
	PluginsFolder        string                       `yaml:"pluginsFolder"`
	Uploaders            []map[string]interface{}     `yaml:"uploaders"`
	InitializedUploaders map[string]uploader.Uploader `yaml:"-"`

	UploaderS3Conf           *UploaderS3Conf           `yaml:"-"`
	UploaderFSConf           *UploaderFSConf           `yaml:"-"`
	UploaderCloudStorageConf *UploaderCloudStorageConf `yaml:"-"`
}
