package handler

import (
	"encoding/json"
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

var (
	products map[string]basket.Item = map[string]basket.Item{
		"VOUCHER": {
			"VOUCHER", "Cabify Voucher", 5.00,
		},
		"TSHIRT": {
			"TSHIRT", "Cabify T-Shirt", 20.00,
		},
		"MUG": {
			"MUG", "Cabify Coffee Mug", 7.50,
		},
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

	product, ok := products[item.Code]

	if !ok {
		Error(w, http.StatusBadRequest, "product not found")

		return
	}

	item.Name = product.Name
	item.Price = product.Price

	bkt := data.(*basket.Basket)
	bkt.Items = append(bkt.Items, item)

	render.JSON(w, http.StatusCreated, bkt)
}
