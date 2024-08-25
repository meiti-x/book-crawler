package cache_storage

import (
	"os"
	"path/filepath"
	"sync"
)

type SafeCache struct {
	dir  string
	lock sync.Mutex
}

func NewSafeCache(dir string) *SafeCache {
	return &SafeCache{dir: dir}
}

func (s *SafeCache) Save(key string, content []byte) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	// Implement your caching logic here
	// Example: Save to a file
	path := filepath.Join(s.dir, key)
	return os.WriteFile(path, content, 0644)
}

func (s *SafeCache) Load(key string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	// Implement your caching load logic here
	path := filepath.Join(s.dir, key)
	return os.ReadFile(path)
}
