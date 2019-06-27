package remotes

import (
	"context"

	"github.com/deps-cloud/rds/api"
	"github.com/deps-cloud/rds/pkg/config"

	"google.golang.org/grpc"
)

var _ Remote = &rdsRemote{}

// NewRDSRemote produces a remote connecting to another RDS
func NewRDSRemote(config *config.Rds) (Remote, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
	}

	cc, err := grpc.Dial(config.Target, opts...)
	if err != nil {
		return nil, err
	}

	return &rdsRemote{
		client: api.NewRepositoryDiscoveryClient(cc),
	}, nil
}

type rdsRemote struct {
	client api.RepositoryDiscoveryClient
}

func (r *rdsRemote) ListRepositories() ([]string, error) {
	response, err := r.client.List(context.Background(), &api.ListRepositoriesRequest{})
	if err != nil {
		return nil, err
	}
	return response.Repositories, nil
}
