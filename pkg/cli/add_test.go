package cli

import (
	"bytes"
	"flag"
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"github.com/gorilla/mux"
	ufcli "github.com/urfave/cli"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCliAdd(t *testing.T) {
	productRepo := product.NewRepo()

	item := basket.Item{"MUG", "Cabify Coffee Mug", 7.50}
	bkt := basket.NewBasket()
	bkt.AddItem(item)

	store := &StorageMock{bkt}

	basketsHandler := handler.NewBaskets(
		store,
		productRepo,
		basket.NewCalculator(),
	)

	router := mux.NewRouter()
	router.HandleFunc("/baskets/{id}/items", basketsHandler.AddItem).Methods(http.MethodPost)

	server := httptest.NewServer(router)

	buf := &bytes.Buffer{}

	app := ufcli.NewApp()
	app.Writer = buf

	flags := flag.NewFlagSet("test", 0)
	flags.String("host", server.URL, "")
	flags.Parse([]string{"1", "MUG"})
	context := ufcli.NewContext(app, flags, nil)

	if err := Add(context); err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(buf.String(), "product added") {
		t.Error("invalid command result")
	}
}
