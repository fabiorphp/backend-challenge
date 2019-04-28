package main

import (
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/gorilla/mux"
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

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods(http.MethodGet)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Fatal(err)
	}
}
