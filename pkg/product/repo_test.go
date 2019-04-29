package product

import (
	"testing"
)

func TestProductRepoGetByCode(t *testing.T) {
	repo := NewRepo()

	if _, err := repo.GetByCode("SHOE"); err == nil {
		t.Errorf(
			"method not returned error: got %v want %v",
			err,
			ErrProductNotFound,
		)
	}

	product, err := repo.GetByCode("MUG")

	if err != nil {
		t.Fatal()
	}

	if product.Code != "MUG" {
		t.Error("invalid product code")
	}
}
