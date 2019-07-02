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

func (s *staticRemote) ListRepositories() ([]string, error) {
	return s.cfg.RepositoryUrls[0:], nil
}
