package uploader

import (
	"mime/multipart"
	"time"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
)

type repository struct {
	uploaders map[string]Uploader
}

var repo repository

func init() {
	repo = repository{
		uploaders: make(map[string]Uploader),
	}
}

// Uploader represents
type Uploader interface {
	Upload(multipart.File, *multipart.FileHeader) error
}

// S3Uploader is an interface extending the Uploader with
// S3 related capabilities
type S3Uploader interface {
	Uploader
	GetS3PresignedPutURL(time.Duration, string, string) (string, error)
}

// RegisterUploader registers an uploader within the uploader repository
// It returns an error if two uploaders are registered under the same name
func RegisterUploader(uploaderName string, u Uploader) (alreadyRegistered error) {
	if repo.uploaders[uploaderName] != nil {
		return errors.NewErrDuplicateUploader(uploaderName)
	}

	repo.uploaders[uploaderName] = u
	return nil
}

// Upload calls the underlying uploader Upload function
func Upload(uploaderName string, file multipart.File, fileHeader *multipart.FileHeader) error {
	return repo.uploaders[uploaderName].Upload(file, fileHeader)
}

// GetS3PresignedPutURL is a function available only when using s3 uploaders
// It returns a presigned url in order to delegate upload on front side only
func GetS3PresignedPutURL(uploaderName string, expire time.Duration, fileName, fileType string) (string, error) {
	s3Uploader, ok := repo.uploaders[uploaderName].(S3Uploader)
	if !ok {
		panic("you must use s3 driver in order to use GetS3PresignedPutURL")
	}

	return s3Uploader.GetS3PresignedPutURL(expire, fileName, fileType)
}
