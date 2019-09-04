package errors

import (
	"fmt"
)

// ErrFileTooLarge is the error triggered if a file is too large for
// uploading. Please see the config models to figure out the max allowed
// size
type ErrFileTooLarge struct {
	limit int64
}

// NewErrFileTooLarge returns an instance of ErrFileTooLarge
func NewErrFileTooLarge(limit int64) ErrFileTooLarge {
	return ErrFileTooLarge{
		limit: limit,
	}
}

func (e ErrFileTooLarge) Error() string {
	return fmt.Sprintf("the file you're trying to upload is too large : the maximum allowed is %v", e.limit)
}
