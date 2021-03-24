package memory

import (
	"sync"
	"time"
)

// Item is a cached reference
type Item struct {
	UserID     int64
	Status     int64
	Expiration int64
}

// Expired returns true if the item has expired.
func (item Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().Unix() > item.Expiration
}

// Storage mechanism for caching strings in memory
type Storage struct {
	items map[string]Item
	mu    *sync.RWMutex
}

// NewStorage creates a new in memory storage
func NewStorage() *Storage {
	return &Storage{
		items: make(map[string]Item),
		mu:    &sync.RWMutex{},
	}
}

// Get a cached content by key
func (s Storage) Get(key string) *Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[key]
	if !ok {
		return nil
	}

	if item.Expired() {
		delete(s.items, key)
		return nil
	}

	return &item
}

// Set a cached content by key
func (s Storage) Set(key string, item Item) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[key] = item
}
