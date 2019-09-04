package errors

import "fmt"

// ErrDuplicateUploader is the error returned in case a
// duplicate uploader is being registered
type ErrDuplicateUploader struct {
	Name string
}

// NewErrDuplicateUploader returns a new NewErrDuplicateUploader
func NewErrDuplicateUploader(uploaderName string) (errDuplicateUploader error) {
	return ErrDuplicateUploader{
		Name: uploaderName,
	}
}

func (e ErrDuplicateUploader) Error() string {
	return fmt.Sprintf("an uploader with the name %v is already registered", e.Name)
}
