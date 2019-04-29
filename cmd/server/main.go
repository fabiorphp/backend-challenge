package main

import (
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
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

	basketsHandler := handler.NewBaskets(storage.NewMemory())

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods(http.MethodGet)
	router.HandleFunc("/baskets", basketsHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/baskets/{id}", basketsHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/baskets/{id}/items", basketsHandler.AddItem).Methods(http.MethodPost)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Fatal(err)
	}
}
