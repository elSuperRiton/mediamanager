package fs

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/elSuperRiton/mediamanager/pkg/models"
)

type Uploader struct {
	conf *models.UploaderFSConf
}

func NewUploader(conf *models.UploaderFSConf) (Uploader, error) {
	return Uploader{
		conf: conf,
	}, nil
}

func (u Uploader) Upload(f multipart.File, fH *multipart.FileHeader) error {

	if _, err := os.Stat(u.conf.Folder); os.IsNotExist(err) {
		os.Mkdir(u.conf.Folder, os.ModePerm)
	}

	// check if plugin exists and extact filename from it
	filename := fH.Filename
	if u.conf.PluginName != "" {
		n, _ := u.conf.GetFileName(fH)
		filename = n
	}

	newFile, err := os.OpenFile(u.conf.Folder+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(newFile, f)
	if err != nil {
		return err
	}

	return nil
}
