package basket

import (
	"github.com/fabiorphp/backend-challenge/pkg/product"
	"testing"
)

func TestBasketCalculator(t *testing.T) {
	repo := product.NewRepo()

	calc := NewCalculator(
		Sum,
		BuyTwoGetOneFree(repo, "VOUCHER"),
		BuyTwoGetOneFree(repo, "SHOE"),
		BulkDiscount(repo, "TSHIRT", 19.0),
		BulkDiscount(repo, "PANTS", 10.0),
	)

	bkt := NewBasket()
	bkt.AddItem(Item{"VOUCHER", "Cabify Voucher", 5.00})
	bkt.AddItem(Item{"VOUCHER", "Cabify Voucher", 5.00})
	bkt.AddItem(Item{"VOUCHER", "Cabify Voucher", 5.00})
	bkt.AddItem(Item{"TSHIRT", "Cabify T-Shirt", 20.00})
	bkt.AddItem(Item{"TSHIRT", "Cabify T-Shirt", 20.00})
	bkt.AddItem(Item{"TSHIRT", "Cabify T-Shirt", 20.00})
	bkt.AddItem(Item{"MUG", "Cabify Coffee Mug", 7.50})

	if result := calc.Calculate(bkt); result != 74.5 {
		t.Errorf("wrong calculation value: got %v want %v", result, 74.5)
	}
}
