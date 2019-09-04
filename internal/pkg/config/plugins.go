package config

import (
	"fmt"
	"mime/multipart"
	"plugin"

	"github.com/elSuperRiton/mediamanager/pkg/errors"
)

type (
	pluginGetFileNameSignature func(file *multipart.FileHeader) (newFileName string, err error)
	// Plugin defines the set of function needed to have a compliant
	// plugin loaded
	Plugin interface {
		GetFileName(file *multipart.FileHeader) (newFileName string, err error)
	}
)

func (c configurableUploader) loadPlugin(pluginFolder string) error {
	// grab uploader plugin
	if c != nil {
		pluginName, ok := c["pluginName"].(string)
		if !ok {
			return fmt.Errorf("plugin must be of type string for uploader %v", c["type"])
		}

		p, err := plugin.Open(pluginFolder + pluginName)
		if err != nil {
			return err
		}

		symbol, err := p.Lookup("GetFileName")
		if err != nil {
			return err
		}

		getNameFunc, ok := symbol.(func(file *multipart.FileHeader) (newFileName string, err error))
		if !ok {
			return errors.ErrWrongGetFileNameImplementation(pluginName)
		}

		c["plugin"] = p
		c["GetFileName"] = getNameFunc
	}

	return nil
}
