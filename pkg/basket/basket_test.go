package basket

import (
	"reflect"
	"testing"
)

func TestNewBasket(t *testing.T) {
	if reflect.TypeOf(NewBasket()) != reflect.TypeOf(&Basket{}) {
		t.Error("Invalid type returned")
	}
}
