package handler

import (
	"bytes"
	"encoding/json"
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	basketsHandler = NewBaskets(storage.NewMemory())
)

type (
	StorageMock struct {
		basket interface{}
	}
)

func (s *StorageMock) Delete(key string) {}

func (s *StorageMock) Fetch(key string) (interface{}, error) {
	return s.basket, nil
}

func (s *StorageMock) Save(key string, value interface{}) {}

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
			http.StatusCreated,
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
			http.StatusNoContent,
		)
	}
}

func TestBasketsAddItemHanderWhenBasketNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/baskets/1/items", nil)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(basketsHandler.AddItem)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusNotFound,
		)
	}
}

func TestBasketsAddItemHanderWithInvalidRequest(t *testing.T) {
	buf := bytes.NewBufferString(`{"code":"}`)

	req, err := http.NewRequest(http.MethodPost, "/baskets/1/items", buf)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	basketsHandler = NewBaskets(new(StorageMock))
	handler := http.HandlerFunc(basketsHandler.AddItem)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestBasketsAddItemHanderWithInvalidProduct(t *testing.T) {
	buf := bytes.NewBufferString(`{"code":"SHOES"}`)

	req, err := http.NewRequest(http.MethodPost, "/baskets/1/items", buf)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	basketsHandler = NewBaskets(new(StorageMock))
	handler := http.HandlerFunc(basketsHandler.AddItem)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusBadRequest,
		)
	}
}

func TestBasketsAddItemHander(t *testing.T) {
	buf := bytes.NewBufferString(`{"code":"MUG"}`)

	req, err := http.NewRequest(http.MethodPost, "/baskets/1/items", buf)

	if err != nil {
		t.Fail()
	}

	rec := httptest.NewRecorder()

	store := &StorageMock{basket.NewBasket()}
	basketsHandler = NewBaskets(store)

	handler := http.HandlerFunc(basketsHandler.AddItem)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf(
			"handler returned invalid status code: got %v want %v",
			status,
			http.StatusCreated,
		)
	}
}
