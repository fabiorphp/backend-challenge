package storage

import (
	"errors"
)

type (
	Storage interface {
		Delete(key string)
		Fetch(key string) (interface{}, error)
		Save(key string, value interface{})
	}
)

var (
	KeyNotFound = errors.New("key not found")
)
