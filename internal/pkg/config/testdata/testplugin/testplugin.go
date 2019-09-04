package main

import (
	"mime/multipart"
)

// GetFileName simply returns the fileName
func GetFileName(file *multipart.FileHeader) (newFileName string, err error) {
	return file.Filename, nil
}
