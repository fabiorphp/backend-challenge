package main

import (
	"log"
	"net/http"
)

var (
	address = "localhost:9000"
	appName = "basket"
	version = "0.0.0"
)

func main() {

	log.Printf("%s service - %s", appName, version)
	log.Printf("starting server on %s", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}
