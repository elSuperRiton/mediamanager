package config

import "log"

// uploaderPathValidator is a map with validation function attached in order
// to ensure uniqueness of path values
type uploaderPathValidator map[string]uint8

// validate makes sure an uploader path has been declared and that it is
// unique
func (u uploaderPathValidator) validate(uploaderPath, uploaderType string) {
	if uploaderPath == "" {
		log.Fatalf("path must be provided for uploader %v", uploaderType)
	}

	u[uploaderPath]++
	if u[uploaderPath] > 1 {
		log.Fatalf("an uploader with a path of \"%v\" already exists", uploaderPath)
	}
}

type configurableUploader map[string]interface{}

func (c configurableUploader) validatePath(validator uploaderPathValidator) configurableUploader {
	// grab uploader path and validate
	path, ok := c["path"].(string)
	if !ok {
		log.Fatalf("path must be of type string for uploader %v", c["type"])
	}

	validator.validate(path, c["type"].(string))
	return c
}
