package remotes

import "github.com/deps-cloud/discovery/pkg/config"

type Repository struct {
	RepositoryURL string
	Clone *config.Clone
}

type FetchRepositoriesRequest struct {

}

type FetchRepositoriesResponse struct {
	Repositories []*Repository
}

type Remote interface {
	FetchRepositories(request *FetchRepositoriesRequest) (*FetchRepositoriesResponse, error)
}
