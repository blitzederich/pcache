// Copyright 2022 Alexander Samorodov <blitzerich@gmail.com>

package store

type Store[K comparable, V any] struct {
	store map[K]V
}

func New[K string, V string | []byte]() *Store[K, V] {
	return &Store[K, V]{
		store: make(map[K]V),
	}
}

func (s *Store[K, V]) Get(key K) (V, bool) {
	value, ok := s.store[key]
	return value, ok
}

func (s *Store[K, V]) Set(key K, value V) {
	s.store[key] = value
}

func (s *Store[K, V]) Update(key K, value V) {
	s.store[key] = value
}

func (s *Store[K, V]) Delete(key K) {
	delete(s.store, key)
}

func (s *Store[K, V]) GetStore() *map[K]V {
	return &s.store
}
