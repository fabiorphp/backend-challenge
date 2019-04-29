package main

import (
	"flag"
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	address string
	appName = "basket"
	version = "0.0.0"
)

func init() {
	flag.StringVar(&address, "listen", "localhost:9000", "Address and port on which App will accept HTTP requests")
}

func main() {
	flag.Parse()

	log.Printf("%s service - %s", appName, version)
	log.Printf("starting server on %s", address)

	basketsHandler := handler.NewBaskets(storage.NewMemory(), product.NewRepo())

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods(http.MethodGet)
	router.HandleFunc("/baskets", basketsHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/baskets/{id}", basketsHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/baskets/{id}/items", basketsHandler.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/baskets/{id}/amount", basketsHandler.Amount).Methods(http.MethodGet)

	err := http.ListenAndServe(address, router)

	if err != nil {
		log.Fatal(err)
	}
}
