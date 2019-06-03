package store

import (
	"github.com/deps-cloud/rds/api"
	"sync"
)

// NewRepositoryStore creates a synchronized store of information.
func NewRepositoryStore() *RepositoryStore {
	return &RepositoryStore{
		repositories: make(map[string]*api.Repository),
		lock:         &sync.Mutex{},
	}
}

// RepositoryStore defines a locked structure for easy sharing across goroutine bounds.
type RepositoryStore struct {
	repositories map[string]*api.Repository
	lock         *sync.Mutex
}

// Insert appends new entires to the internal storage structure.
func (s *RepositoryStore) Insert(repositories []*api.Repository) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, repository := range repositories {
		s.repositories[repository.Url] = repository
	}
}

// Snapshot takes a snapshot of the datastore, returning the repositories it knows about.
func (s *RepositoryStore) Snapshot() []*api.Repository {
	s.lock.Lock()
	defer s.lock.Unlock()

	repositories := make([]*api.Repository, 0)

	for _, repository := range s.repositories {
		repositories = append(repositories, repository)
	}

	return repositories
}
