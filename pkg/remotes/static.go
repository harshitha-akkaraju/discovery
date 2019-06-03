package remotes

import (
	"github.com/deps-cloud/rds/api"
	"github.com/deps-cloud/rds/pkg/config"
)

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

func (s *staticRemote) ListRepositories() ([]*api.Repository, error) {
	urls := s.cfg.GetRepositoryUrls()
	repositories := make([]*api.Repository, len(urls))
	for i, url := range urls {
		repositories[i] = &api.Repository{
			Url: url,
		}
	}
	return repositories, nil
}
