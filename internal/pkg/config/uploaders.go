package config

import (
	"fmt"

	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/uploader"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/fs"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/s3"
)

func (c configurableUploader) initUploader(conf *models.MediaManagerConfig) error {

	var (
		u   uploader.Uploader
		err error
	)

	switch c["type"] {
	// S3 UPLOADER
	// Register an Amazon s3 uploader
	case models.UploaderTypeS3:
		var s3UploaderConf models.UploaderS3Conf
		mapStructConf(c, &s3UploaderConf)
		u, err = s3.NewUploader(s3UploaderConf)
	// FILESYSTEM UPLOADER
	// Register an `classic` file system uploader
	case models.UploaderTypeFileSystem:
		var fsUploaderConf models.UploaderFSConf
		mapStructConf(c, &fsUploaderConf)
		u, err = fs.NewUploader(&fsUploaderConf)
	// GOOGLE CLOUD STORAGE
	// Register a google cloud storage uploader
	case models.UploaderTypeCloudStorage:
		var csUploader models.UploaderCloudStorageConf
		mapStructConf(c, &csUploader)
		conf.UploaderCloudStorageConf = &csUploader
	// UNSUPPORTED UPLOADER TYPE
	default:
		return fmt.Errorf("unsuported uploader of type : %v", c["type"])
	}

	if err != nil {
		return err
	}

	uploader.RegisterUploader(c["path"].(string), u)
	return nil
}
