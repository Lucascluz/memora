package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	store map[string][]byte
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string][]byte),
		mu:    sync.Mutex{},
	}
}

func (c *Cache) Set(key string, value []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if value is nil
	if value == nil {
		return errors.New("cannot insert null value")
	}

	//set value (overrides if key already exists)
	c.store[key] = value

	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if exists
	value, ok := c.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	return value, nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if exists
	_, ok := c.store[key]
	if !ok {
		return errors.New("key not found")
	}

	// delete entry
	delete(c.store, key)

	return nil
}
