package main

import "errors"

// GetFileName is a wrong implementation of a GetFileName from
// an uploader plugin for testing purpose
func GetFileName(wrongArgument bool) (fakeError error) {
	return errors.New("fake error")
}

// PostUpload is a wrong implementation of a PostUpload from
// an uploader plugin for testing purpose
func PostUpload(wrongArgument bool) (fakeError error) {
	return errors.New("fake error")
}
