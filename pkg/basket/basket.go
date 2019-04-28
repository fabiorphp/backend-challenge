package basket

import (
	"time"
)

type (
	Basket struct {
		ID int64 `json:"id"`
	}
)

func NewBasket() *Basket {
	return &Basket{
		ID: time.Now().Unix(),
	}
}
