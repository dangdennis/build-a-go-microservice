package service

import (
	"log"
	"net/http"
)

// StartWebServer initiates our http server
func StartWebServer(port string) {
	r := NewRoute()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
