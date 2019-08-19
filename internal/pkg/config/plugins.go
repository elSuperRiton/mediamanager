package config

import (
	"fmt"
	"log"
	"mime/multipart"
	"plugin"
)

func (c configurableUploader) loadPlugin() configurableUploader {
	// grab uploader plugin
	if c["pluginName"] != nil {
		pluginName, ok := c["pluginName"].(string)
		fmt.Println(Conf.PluginsFolder)
		if !ok {
			log.Fatalf("plugin must be of type string for uploader %v", c["type"])
		}

		p, err := plugin.Open(Conf.PluginsFolder + pluginName)
		if err != nil {
			panic(err)
		}

		symbol, err := p.Lookup("GetFileName")
		if err != nil {
			panic(err)
		}

		getNameFunc, ok := symbol.(func(file *multipart.FileHeader) (newFileName string, err error))
		if !ok {
			panic("Plugin has no 'func(file *multipart.FileHeader) (newFileName string, err error)' function")
		}

		c["plugin"] = p
		c["GetFileName"] = getNameFunc
	}

	return c
}
