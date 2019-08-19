package uploader

import (
	"mime/multipart"
	"time"
)

// Uploader is the package's main interface defining methods
// available in all drivers
type Uploader interface {
	Upload(multipart.File, *multipart.FileHeader) error
}

// S3Uploader is an interface extending the Uploader with
// S3 related capabilities
type S3Uploader interface {
	Uploader
	GetS3PresignedPutURL(time.Duration, string, string) (string, error)
}

var repository Uploader

// InitUploader takes in an Uploader and initializes the uploader
// repository with the provided uploader
func InitUploader(u Uploader) {
	repository = u
}

// Upload calls Upload method of the underlying package Uploader
// interface set by InitUploader
func Upload(f multipart.File, fH *multipart.FileHeader) error {
	return repository.Upload(f, fH)
}

// GetS3PresignedPutURL is a method available when using the package's
// S3 driver
// A panic will be triggered if this function is called under usage of
// any other driver
func GetS3PresignedPutURL(expire time.Duration, fileName, fileType string) (string, error) {
	s3Uploader, ok := repository.(S3Uploader)
	if !ok {
		panic("you must use s3 driver in order to use GetS3PresignedPutURL")
	}

	return s3Uploader.GetS3PresignedPutURL(expire, fileName, fileType)
}
