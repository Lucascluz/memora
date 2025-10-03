package cache

import (
	"errors"
	"sync"
	"time"
)

type entry struct {
	value []byte
	ttl   int64
}

type Cache struct {
	store map[string]entry
	mu    sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]entry),
		mu:    sync.Mutex{},
	}
}

func (c *Cache) Set(key string, value []byte, ttl int64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if value is nil
	if value == nil {
		return errors.New("cannot insert null value")
	}

	// check if ttl is >= 0
	if ttl < 0 {
		return errors.New("cannot insert expired entry")
	}

	//set value (overrides if key already exists)
	c.store[key] = entry{value: value, ttl: ttl}

	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// check if exists
	entry, ok := c.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}

	// check if expired
	if entry.ttl < time.Now().Unix() {
		return nil, errors.New("entry expired")
	}

	return entry.value, nil
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
