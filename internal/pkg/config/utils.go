package config

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

// mapStructConf is a helper function that panics if there an error decoding a
// structure using mapstructure
func mapStructConf(input, output interface{}) {
	if err := mapstructure.Decode(input, output); err != nil {
		log.Fatalf("error mapping structure while building configuration : %v", err)
	}
}
