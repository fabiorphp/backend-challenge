package storage

import (
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	store := NewMemory()

	store.Save("a", "b")

	if data, _ := store.Fetch("a"); data != "b" {
		t.Errorf(
			"storage returned invalid data: got %v want %v",
			data,
			"b",
		)
	}

	store.Delete("a")

	if _, err := store.Fetch("a"); err != KeyNotFound {
		t.Error("Invalid error message")
	}
}
