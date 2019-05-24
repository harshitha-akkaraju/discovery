package service

import (
	"time"

	"github.com/deps-cloud/rds/api"
	"github.com/deps-cloud/rds/pkg/remotes"
	"github.com/deps-cloud/rds/pkg/service/store"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func load(remote remotes.Remote, repositoryStore *store.RepositoryStore) {
	logrus.Infof("[service.rds] loading repositories")

	if repositories, err := remote.ListRepositories(); err != nil {
		logrus.Errorf("[service.rds] failed to load repositories: %v", err)
	} else {
		repositoryStore.Insert(repositories)
	}
}

// NewServer creates an rds server using the specified remote
func NewServer(remote remotes.Remote) api.RepositoryDiscoveryServiceServer {
	repositoryStore := store.NewRepositoryStore()

	go func() {
		load(remote, repositoryStore)
		for range time.Tick(30 * time.Minute) {
			load(remote, repositoryStore)
		}
	}()

	return &repositoryDiscoveryService{
		repositoryStore: repositoryStore,
	}
}

var _ api.RepositoryDiscoveryServiceServer = &repositoryDiscoveryService{}

type repositoryDiscoveryService struct {
	repositoryStore *store.RepositoryStore
}

func (s *repositoryDiscoveryService) List(ctx context.Context, req *api.ListRepositoriesRequest) (*api.ListRepositoriesResponse, error) {
	repositories := s.repositoryStore.Snapshot()

	return &api.ListRepositoriesResponse{
		Repositories: repositories,
	}, nil
}
