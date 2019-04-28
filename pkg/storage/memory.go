package storage

import (
	"sync"
)

type (
	Memory struct {
		data sync.Map
	}
)

func NewMemory() *Memory {
	return &Memory{
		data: sync.Map{},
	}
}

func (m *Memory) Delete(key string) {
	m.data.Delete(key)
}

func (m *Memory) Fetch(key string) (interface{}, error) {
	v, ok := m.data.Load(key)

	if !ok {
		return nil, KeyNotFound
	}

	return v, nil
}

func (m *Memory) Save(key string, value interface{}) {
	m.data.Store(key, value)
}
