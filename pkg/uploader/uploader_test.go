package uploader

import (
	"mime/multipart"
	"testing"
	"time"
)

type mockGenericUploader struct{}

func (m mockGenericUploader) Upload(f multipart.File, fH *multipart.FileHeader) error {
	return nil
}

type mockS3Uploader struct{}

func (m mockS3Uploader) Upload(f multipart.File, fH *multipart.FileHeader) error {
	return nil
}

func (m mockS3Uploader) GetS3PresignedPutURL(expire time.Duration, fileName, fileType string) (string, error) {
	return "fakesigneduri", nil
}

func Test_InitUploader(t *testing.T) {
	t.Error("Not implemented yet")
}

func Test_Upload(t *testing.T) {
	t.Error("Not implemented yet")
}

func Test_GetS3PresignedPutURL(t *testing.T) {
	t.Error("Not implemented yet")
}
