package main

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

// GetFileName returns a filename with a unix timestamp
func GetFileName(file *multipart.FileHeader) (newFileName string, err error) {
	return strconv.Itoa(int(time.Now().Unix())) + file.Filename, nil
}

// PostUpload is executed after a successfull upload
func PostUpload() error {
	fmt.Println("postupload")
	return nil
}
