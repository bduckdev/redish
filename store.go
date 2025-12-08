package main

import (
	"sync"
)

// Store is a thread-safe key-value store.
// It uses a RWMutex to allow multiple readers (GET) at once,
// but locks everything for a single writer (SET).
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStore initializes the map
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Set acquires a Write Lock (blocking everyone else) to ensure safety.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	// TODO: Trigger Write-Ahead Log (WAL) append here
}

// Get acquires a Read Lock. Multiple threads can Get at the same time.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]
	return val, ok
}
