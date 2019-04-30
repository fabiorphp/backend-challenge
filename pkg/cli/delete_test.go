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

func TestCliDelete(t *testing.T) {
	productRepo := product.NewRepo()

	basketsHandler := handler.NewBaskets(
		storage.NewMemory(),
		productRepo,
		basket.NewCalculator(),
	)

	router := mux.NewRouter()
	router.HandleFunc("/baskets/{id}", basketsHandler.Delete).Methods(http.MethodDelete)

	server := httptest.NewServer(router)

	buf := &bytes.Buffer{}

	app := ufcli.NewApp()
	app.Writer = buf

	flags := flag.NewFlagSet("test", 0)
	flags.String("host", server.URL, "")
	flags.Parse([]string{"1"})
	context := ufcli.NewContext(app, flags, nil)

	if err := Delete(context); err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(buf.String(), "basket deleted") {
		t.Error("invalid command result")
	}
}
