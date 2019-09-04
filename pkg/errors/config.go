package errors

import "fmt"

// ErrWrongGetFileNameImplementation formats an error for a wrong GetFileName implementation
// when loading a plugin
func ErrWrongGetFileNameImplementation(pluginName string) error {
	return fmt.Errorf("plugin %v has no 'func(file *multipart.FileHeader) (newFileName string, err error)' function", pluginName)
}

// ConfigurableUploaderPathErr returns a new ConfigurableUploaderPathErr
func ConfigurableUploaderPathErr(path string) error {
	return fmt.Errorf("uplaoder with a path of \"%v\" already exists", path)
}
