package cli

import (
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

func Serve(c *cli.Context) error {
	log.Printf("%s service - %s", c.App.Name, c.App.Version)
	log.Printf("starting server on %s", c.String("listen"))

	productRepo := product.NewRepo()

	basketsHandler := handler.NewBaskets(
		storage.NewMemory(),
		productRepo,
		basket.NewCalculator(
			basket.Sum,
			basket.BuyTwoGetOneFree(productRepo, "VOUCHER"),
			basket.BulkDiscount(productRepo, "TSHIRT", 19.0),
		),
	)

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods(http.MethodGet)
	router.HandleFunc("/baskets", basketsHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/baskets/{id}", basketsHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/baskets/{id}/items", basketsHandler.AddItem).Methods(http.MethodPost)
	router.HandleFunc("/baskets/{id}/amount", basketsHandler.Amount).Methods(http.MethodGet)

	return http.ListenAndServe(c.String("listen"), router)
}
