package handler

import (
	"encoding/json"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	basketsHandler = NewBaskets(storage.NewMemory())
)

func TestBasketsCreateHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/baskets", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(basketsHandler.Create)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}

	res := make(map[string]int64)

	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fail()
	}

	if _, ok := res["id"]; !ok {
		t.Error("response id key not found")
	}
}

func TestBasketsDeleteHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/baskets", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(basketsHandler.Delete)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNoContent {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusOK,
		)
	}
}
