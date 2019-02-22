package store

import "sync"

// NewRepositoryStore creates a synchronized store of information.
func NewRepositoryStore() *RepositoryStore {
	return &RepositoryStore{
		repositories: make(map[string]bool),
		lock: &sync.Mutex{},
	}
}

// RepositoryStore defines a locked structure for easy sharing across goroutine bounds.
type RepositoryStore struct {
	repositories map[string]bool
	lock *sync.Mutex
}

// Insert appends new entires to the internal storage structure.
func (s *RepositoryStore) Insert(repositories []string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, repository := range repositories {
		s.repositories[repository] = true
	}
}

// Snapshot takes a snapshot of the datastore, returning the repositories it knows about.
func (s *RepositoryStore) Snapshot() []string {
	s.lock.Lock()
	defer s.lock.Unlock()

	repositories := make([]string, 0)

	for key := range s.repositories {
		repositories = append(repositories, key)
	}

	return repositories
}