package config

import (
	"testing"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
)

const testDataLocation = "./testdata/"

func TestLoadPlugin(t *testing.T) {
	testSuite := []struct {
		uploader configurableUploader
		expected error
	}{
		{
			uploader: configurableUploader{
				"pluginName": "testplugin.so",
			},
			expected: nil,
		},
		{
			uploader: configurableUploader{
				"pluginName": "testwrongplugin.so",
			},
			expected: errors.ErrWrongGetFileNameImplementation("testwrongplugin.so"),
		},
	}

	for _, test := range testSuite {
		err := test.uploader.loadPlugin(testDataLocation)
		if err != nil && err.Error() != test.expected.Error() {
			t.Errorf("expected error to be \"%v\", got %v", test.expected, err)
		}
	}
}
