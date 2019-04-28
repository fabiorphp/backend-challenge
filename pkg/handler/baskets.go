package handler

import (
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/render"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type (
	Baskets struct {
		store storage.Storage
	}
)

func NewBaskets(store storage.Storage) *Baskets {
	return &Baskets{store}
}

func (b *Baskets) Create(w http.ResponseWriter, r *http.Request) {
	bkt := basket.NewBasket()
	b.store.Save(strconv.FormatInt(bkt.ID, 10), bkt)

	render.JSON(w, http.StatusCreated, bkt)
}

func (b *Baskets) Delete(w http.ResponseWriter, r *http.Request) {
	b.store.Delete(mux.Vars(r)["id"])

	render.JSON(w, http.StatusNoContent, nil)
}
