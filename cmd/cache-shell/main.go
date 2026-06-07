package main

import (
	"sync"
)

type SafeMap struct{
	mu sync.Mutex
	m map[string]string
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock()
	s.m[key] = value
	s.mu.Unlock()
}

func (s *SafeMap) Get(key string) string {
	s.mu.Lock()
	value := s.m[key]
	s.mu.Unlock()
	return value
}