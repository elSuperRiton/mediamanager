// Package config implements formatted I/O with functions analogous to C's printf and scanf.
// The format 'verbs' are derived from C's but are simpler.
package config

import (
	"io/ioutil"
	"log"

	"github.com/elSuperRiton/mediamanager/pkg/models"
	"github.com/elSuperRiton/mediamanager/pkg/uploader"

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
	Conf.InitializedUploaders = make(map[string]uploader.Uploader, len(Conf.Uploaders))

	for _, uploader := range Conf.Uploaders {
		c := configurableUploader(uploader)
		c.validatePath(pathValidator).loadPlugin().initUploader()
	}

	if err := validate.Struct(Conf); err != nil {
		log.Fatal(err)
	}

	return Conf, nil
}
