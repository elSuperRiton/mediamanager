package uploader

import (
	"mime/multipart"
	"testing"
	"time"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/fs"
	"github.com/elSuperRiton/mediamanager/pkg/uploader/drivers/s3"
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

func TestRegisterUploader(t *testing.T) {

	s3Uploader, _ := s3.NewUploader(models.UploaderS3Conf{})
	fsUploader, _ := fs.NewUploader(&models.UploaderFSConf{})

	testSuite := []struct {
		expectedErrsNum int
		expectedErrs    []error
		uploaders       []struct {
			name     string
			uploader Uploader
		}
	}{
		{
			expectedErrsNum: 0,
			expectedErrs:    []error{},
			uploaders: []struct {
				name     string
				uploader Uploader
			}{
				{
					name:     "s3",
					uploader: s3Uploader,
				},
				{
					name:     "fs",
					uploader: fsUploader,
				},
			},
		},
		{
			expectedErrsNum: 1,
			expectedErrs: []error{
				errors.NewErrDuplicateUploader("s3"),
			},
			uploaders: []struct {
				name     string
				uploader Uploader
			}{
				{
					name:     "s3",
					uploader: s3Uploader,
				},
				{
					name:     "fs",
					uploader: fsUploader,
				},
				{
					name:     "s3",
					uploader: s3Uploader,
				},
			},
		},
		{
			expectedErrsNum: 2,
			expectedErrs: []error{
				errors.NewErrDuplicateUploader("s3"),
				errors.NewErrDuplicateUploader("fs"),
			},
			uploaders: []struct {
				name     string
				uploader Uploader
			}{
				{
					name:     "s3",
					uploader: s3Uploader,
				},
				{
					name:     "fs",
					uploader: fsUploader,
				},
				{
					name:     "s3",
					uploader: s3Uploader,
				},
				{
					name:     "fs",
					uploader: fsUploader,
				},
			},
		},
	}

	for _, test := range testSuite {
		func() {

			defer func() {
				repo.uploaders = make(map[string]Uploader)
			}()

			errors := []error{}
			for _, uploader := range test.uploaders {
				if err := RegisterUploader(uploader.name, uploader.uploader); err != nil {
					errors = append(errors, err)
				}
			}

			if len(errors) != test.expectedErrsNum {
				t.Errorf("expected number of errors to be \"%v\", got \"%v\"", test.expectedErrsNum, len(errors))
			}

			if len(errors) > 0 {
				for index, err := range errors {
					expectedErr := test.expectedErrs[index]
					if err != expectedErr {
						t.Errorf("expected err to be \"%v\", got \"%v\"", expectedErr, err)
					}
				}
			}
		}()
	}
}

func Test_Upload(t *testing.T) {
	t.Error("Not implemented yet")
}

func Test_GetS3PresignedPutURL(t *testing.T) {
	t.Error("Not implemented yet")
}
