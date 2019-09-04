// Package config implements formatted I/O with functions analogous to C's printf and scanf.
// The format 'verbs' are derived from C's but are simpler.
package config

import (
	"context"
	"io/ioutil"
	"sync"

	"github.com/elSuperRiton/mediamanager/pkg/models"

	"gopkg.in/go-playground/validator.v9"
	yaml "gopkg.in/yaml.v2"
)

var (
	// Conf holds the configuration for media manager
	Conf     *models.MediaManagerConfig
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

// NewConfig reads a yaml configuration file compliant with the uploader
// configuration model
func NewConfig(fileLocation []byte) (uploaderConf *models.MediaManagerConfig, err error) {
	Conf = &models.MediaManagerConfig{}

	// Load configuration file
	yamlFile, err := ioutil.ReadFile(string(fileLocation))
	if err != nil {
		return nil, err
	}

	// Parse configuration
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		return nil, err
	}

	// Validate configuration
	pathValidator := uploaderPathValidator(make(map[string]uint8))
	var wg sync.WaitGroup
	errorChan := make(chan error, 1)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Validatate and initialize uploaders
	wg.Add(len(Conf.Uploaders))
	for _, uploader := range Conf.Uploaders {
		go func(uploader map[string]interface{}) {
			defer wg.Done()

			c := configurableUploader(uploader)

			if err := c.validatePath(pathValidator); err != nil {
				errorChan <- err
				cancel()
			}
			if err := c.loadPlugin(Conf.PluginsFolder); err != nil {
				errorChan <- err
				cancel()
			}
			if err := c.initUploader(Conf); err != nil {
				errorChan <- err
				cancel()
			}
		}(uploader)
	}

	wg.Wait()

	if err := validate.Struct(Conf); err != nil {
		return Conf, err
	}

	return Conf, nil
}
