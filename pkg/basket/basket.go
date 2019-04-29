package basket

import (
	"time"
)

type (
	Basket struct {
		ID    int64  `json:"id"`
		Items []Item `json:"items,omitempty"`
	}

	Item struct {
		Code  string  `json:"code"`
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
)

func NewBasket() *Basket {
	return &Basket{
		ID: time.Now().Unix(),
	}
}

func (b *Basket) AddItem(item Item) {
	b.Items = append(b.Items, item)
}
