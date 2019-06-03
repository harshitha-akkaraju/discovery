package remotes

import "github.com/deps-cloud/rds/api"

// Remote defines a remote source of repositories.
type Remote interface {
	// ListRepositories returns a list of git ssh urls for the given remote interface.
	// The urls will be used to clone repositories on the fly as users navigate the filesystem.
	ListRepositories() ([]*api.Repository, error)
}
