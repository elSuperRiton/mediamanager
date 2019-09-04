package config

import (
	"testing"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
)

func TestValidatePath(t *testing.T) {
	testSuite := []struct {
		uploaders []configurableUploader
		expected  error
	}{
		{
			uploaders: []configurableUploader{
				{
					"type": "s3",
					"path": "s3",
				},
				{
					"type": "fs",
					"path": "s3",
				},
				{
					"type": "cloudplatform",
					"path": "s3",
				},
			},
			expected: errors.ConfigurableUploaderPathErr("s3"),
		},
		{
			uploaders: []configurableUploader{
				{
					"type": "s3",
					"path": "s3",
				},
				{
					"type": "fs",
					"path": "fs",
				},
				{
					"type": "cloudplatform",
					"path": "cloudplatform",
				},
			},
			expected: nil,
		},
		{
			uploaders: []configurableUploader{
				{
					"type": "s3",
					"path": "fs",
				},
				{
					"type": "fs",
					"path": "fs",
				},
				{
					"type": "cloudplatform",
					"path": "cloudplaform",
				},
			},
			expected: errors.ConfigurableUploaderPathErr("fs"),
		},
	}

	for _, test := range testSuite {
		u := uploaderPathValidator{}
		for _, uploader := range test.uploaders {
			err := uploader.validatePath(u)
			if err != nil && err.Error() != test.expected.Error() {
				t.Errorf("expected error to be \"%v\", got \"%v\"", test.expected, err)
			}
		}
	}
}
