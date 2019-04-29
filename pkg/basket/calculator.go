package basket

import (
	"github.com/fabiorphp/backend-challenge/pkg/product"
)

type (
	operation = func(amount float64, items []Item) float64

	Calculator struct {
		operations []operation
	}
)

func NewCalculator(operations ...operation) Calculator {
	return Calculator{operations}
}

func (c *Calculator) Calculate(basket *Basket) float64 {
	var amount float64

	for _, op := range c.operations {
		amount = op(amount, basket.Items)
	}

	return amount
}

func Sum(amount float64, items []Item) float64 {
	for _, i := range items {
		amount = amount + i.Price
	}

	return amount
}

func BuyTwoGetOneFree(productRepo product.Repo, code string) operation {
	return func(amount float64, items []Item) float64 {
		product, err := productRepo.GetByCode(code)

		if err != nil {
			return amount
		}

		var total int

		for _, i := range items {
			if i.Code == product.Code {
				total = total + 1
			}
		}

		if total >= 2 {
			amount = amount - product.Price
		}

		return amount
	}
}

func BulkDiscount(productRepo product.Repo, code string, price float64) operation {
	return func(amount float64, items []Item) float64 {
		product, err := productRepo.GetByCode(code)

		if err != nil {
			return amount
		}

		var total int

		for _, i := range items {
			if i.Code == product.Code {
				total = total + 1
			}
		}

		if total >= 3 {
			amount = amount - (product.Price * float64(total))
			amount = amount + (price * float64(total))
		}

		return amount
	}
}
