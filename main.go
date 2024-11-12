package main

import (
	"log"

	"github.com/pr194/Image-processing-service/service"
)

func main() {
	// initializing the service 
	svc, err := service.New()
	if err != nil {
		log.Fatal("Error initializing service:", err)
	}

	// Starting the server
	svc.RunServer()
}
