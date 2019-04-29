package product

import (
	"errors"
)

type (
	Repo interface {
		GetByCode(code string) (*Product, error)
	}

	Memory struct {
		products map[string]*Product
	}
)

var (
	ErrProductNotFound = errors.New("product not found")
)

func NewRepo() *Memory {
	products := map[string]*Product{
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

	return &Memory{products}
}

func (m *Memory) GetByCode(code string) (*Product, error) {
	product, ok := m.products[code]

	if !ok {
		return nil, ErrProductNotFound
	}

	return product, nil
}
