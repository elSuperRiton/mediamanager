package config

import (
	"log"

	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/fs"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/s3"
)

func (c configurableUploader) initUploader() configurableUploader {
	// load configuration with initialzed uploaders
	// according to their provided types
	switch c["type"] {
	// Create and append s3 uploader
	case models.UploaderTypeS3:
		var s3UploaderConf models.UploaderS3Conf
		mapStructConf(c, &s3UploaderConf)
		s3Uploader, err := s3.NewUploader(s3UploaderConf)
		if err != nil {
			log.Fatal(err)
		}
		Conf.InitializedUploaders[c["path"].(string)] = s3Uploader
	// Create and append file systemp uploader
	case models.UploaderTypeFileSystem:
		var fsUploaderConf models.UploaderFSConf
		mapStructConf(c, &fsUploaderConf)
		fsUploader, err := fs.NewUploader(&fsUploaderConf)
		if err != nil {
			log.Fatal(err)
		}
		Conf.InitializedUploaders[c["path"].(string)] = fsUploader
	// Create and append cloud storage  uploader
	case models.UploaderTypeCloudStorage:
		var csUploader models.UploaderCloudStorageConf
		mapStructConf(c, &csUploader)
		Conf.UploaderCloudStorageConf = &csUploader
	default:
		log.Fatalf("unsuported uploader of type : %v", c["type"])
	}

	return c
}
