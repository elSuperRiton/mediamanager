package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/elSuperRiton/mediamanager/internal/api/router"
	"github.com/elSuperRiton/mediamanager/internal/pkg/config"
	"github.com/elSuperRiton/mediamanager/pkg/models"
)

var (
	conf                      *models.MediaManagerConfig
	configurationFileLocation string
)

func init() {
	flag.StringVar(&configurationFileLocation, "config", "./config.yaml", "location of a yaml configuration file")
	flag.Parse()

	// loadConfiguration
	loadConf()
}

func main() {
	// loadConf()

	srv := &http.Server{
		Addr:        config.Conf.Port,
		ReadTimeout: config.Conf.TimeOut,
		Handler:     router.New(conf),
	}

	// Gracefull shutdown
	//
	// We listen to SIGTERM / SIGINT / Interrupt signal and if received we
	// call the shutdown signal to the server passing in an empty context in
	// order to handle all iddle connection before closing the server
	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
		signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGINT)  // sigterm signal sent from kubernetes

		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v\n", err)
		}
	}()

	log.Printf("Starging HTTP server on port %v\n", config.Conf.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Printf("HTTP server ListenAndServe: %v\n", err)
	}
}

// loadConfiguration reads the configuration file with the location set
// in the configurationFileLocation and parses it into an UploaderConf struct
func loadConf() {
	var err error
	if conf, err = config.NewConfig([]byte(configurationFileLocation)); err != nil {
		log.Fatalf("error loading configuration file : %v", err)
	}
}
