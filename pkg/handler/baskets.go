package handler

import (
	"encoding/json"
	"github.com/fabiorphp/backend-challenge/pkg/basket"
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"github.com/fabiorphp/backend-challenge/pkg/render"
	"github.com/fabiorphp/backend-challenge/pkg/storage"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type (
	Baskets struct {
		store       storage.Storage
		productRepo product.Repo
		calc        basket.Calculator
	}
)

func NewBaskets(store storage.Storage, productRepo product.Repo, calc basket.Calculator) *Baskets {
	return &Baskets{store, productRepo, calc}
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

func (b *Baskets) AddItem(w http.ResponseWriter, r *http.Request) {
	data, err := b.store.Fetch(mux.Vars(r)["id"])

	if err != nil {
		Error(w, http.StatusNotFound, "basket not found")

		return
	}

	var item basket.Item

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		Error(w, http.StatusBadRequest, "invalid json request")

		return
	}

	product, err := b.productRepo.GetByCode(item.Code)

	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())

		return
	}

	item.Name = product.Name
	item.Price = product.Price

	bkt := data.(*basket.Basket)
	bkt.AddItem(item)

	render.JSON(w, http.StatusCreated, bkt)
}

func (b *Baskets) Amount(w http.ResponseWriter, r *http.Request) {
	data, err := b.store.Fetch(mux.Vars(r)["id"])

	if err != nil {
		Error(w, http.StatusNotFound, "basket not found")

		return
	}

	bkt := data.(*basket.Basket)

	amount := b.calc.Calculate(bkt)

	payload := map[string]float64{"amount": amount}

	render.JSON(w, http.StatusOK, payload)
}
