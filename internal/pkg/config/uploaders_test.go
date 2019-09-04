package config

import (
	"testing"

	"github.com/elSuperRiton/mediamanager/pkg/models"
)

func TestInitUploaders(t *testing.T) {
	testSuite := []struct {
		uploader               configurableUploader
		initalizedUploaderType interface{}
	}{
		{
			initalizedUploaderType: models.UploaderCloudStorageConf{},
			uploader: map[string]interface{}{
				"type": models.UploaderTypeS3,
				"path": "s3",
			},
		},
		{
			initalizedUploaderType: models.UploaderFSConf{},
			uploader: map[string]interface{}{
				"type": models.UploaderTypeS3,
				"path": "fs",
			},
		},
		{
			initalizedUploaderType: models.UploaderS3Conf{},
			uploader: map[string]interface{}{
				"type": models.UploaderTypeS3,
				"path": "cloudstorage",
			},
		},
	}

	for _, test := range testSuite {
		var conf *models.MediaManagerConfig
		err := test.uploader.initUploader(conf)
		if err != nil {
			t.Error(err)
		}
	}
}
