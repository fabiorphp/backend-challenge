package cli

import (
	"errors"
)

var (
	ErrTooManyArgs    = errors.New("too many arguments")
	ErrBasketID       = errors.New("basket id is required")
	ErrBasketNotFound = errors.New("basket not found")
)
