package remotes

import "github.com/deps-cloud/discovery/pkg/config"

var _ Remote = &staticRemote{}

// NewStaticRemote produces a new remote from static configuration
func NewStaticRemote(cfg *config.Static) Remote {
	return &staticRemote{
		cfg: cfg,
	}
}

type staticRemote struct {
	cfg *config.Static
}

func (s *staticRemote) FetchRepositories(request *FetchRepositoriesRequest) (*FetchRepositoriesResponse, error) {
	repositories := make([]*Repository, len(s.cfg.RepositoryUrls))
	for i, repositoryURL := range s.cfg.RepositoryUrls {
		repositories[i] = &Repository{
			RepositoryURL: repositoryURL,
			Clone:         s.cfg.GetClone(),
		}
	}
	return &FetchRepositoriesResponse{
		Repositories: repositories,
	}, nil
}
