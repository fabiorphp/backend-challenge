package cli

import (
	"bytes"
	"flag"
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/handler"
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"github.com/gorilla/mux"
	ufcli "github.com/urfave/cli"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCliCreate(t *testing.T) {
	productRepo := product.NewRepo()

	basketsHandler := handler.NewBaskets(
		storage.NewMemory(),
		productRepo,
		basket.NewCalculator(),
	)

	router := mux.NewRouter()
	router.HandleFunc("/baskets", basketsHandler.Create).Methods(http.MethodPost)

	server := httptest.NewServer(router)

	buf := &bytes.Buffer{}

	app := ufcli.NewApp()
	app.Writer = buf

	flags := flag.NewFlagSet("test", 0)
	flags.String("host", server.URL, "")
	context := ufcli.NewContext(app, flags, nil)

	if err := Create(context); err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(buf.String(), "basket id") {
		t.Error("invalid command result")
	}
}
