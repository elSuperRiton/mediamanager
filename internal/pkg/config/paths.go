package config

import (
	"fmt"
	"sync"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
)

var uploaderPathValidatorMtx = sync.RWMutex{}

type (
	// uploaderPathValidator is a map with validation function attached in order
	// to ensure uniqueness of path values
	uploaderPathValidator map[string]uint8
	// configurableUploader is a helper type for validating a series of uploaders.
	// It wraps a map[string]interface{} awaiting at least a "type" value and a "path"
	// value
	configurableUploader map[string]interface{}
)

func (c configurableUploader) validatePath(validator uploaderPathValidator) error {
	uploaderPathValidatorMtx.Lock()
	defer uploaderPathValidatorMtx.Unlock()

	// grab uploader path and validate
	path, ok := c["path"].(string)
	if !ok {
		return fmt.Errorf("path must be of type string for uploader %v", c["type"])
	}

	if path == "" {
		return fmt.Errorf("path must be defined for uploader %v", c["type"])
	}

	validator[path]++

	if validator[path] > 1 {
		return errors.ConfigurableUploaderPathErr(path)
	}

	return nil
}
