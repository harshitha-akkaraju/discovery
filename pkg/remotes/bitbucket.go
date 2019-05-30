package remotes

import (
	"fmt"

	"github.com/davidji99/go-bitbucket"
	"github.com/deps-cloud/rds/pkg/config"
	"github.com/sirupsen/logrus"
)

// NewBitbucketRemote constructs a new remote implementation that speaks with Bitbucket
// for repository related information.
func NewBitbucketRemote(cfg *config.Bitbucket) (Remote, error) {
	var client *bitbucket.Client

	if basic := cfg.GetBasic(); basic != nil {
		username := basic.Username
		password := ""
		if pwd := basic.GetPassword(); pwd != nil {
			password = pwd.Value
		}

		client = bitbucket.NewBasicAuth(username, password)
	} else {
		return nil, fmt.Errorf("auth format not supported")
	}

	return &bitbucketRemote{
		client: client,
		config: cfg,
	}, nil
}

var _ Remote = &bitbucketRemote{}

func convertRepositoriesResponse(response interface{}, strategy string) []string {
	rdata := response.(map[string]interface{})
	values := rdata["values"].([]interface{})

	repos := make([]string, 0, len(values))
	for _, value := range values {
		val := value.(map[string]interface{})

		links := val["links"].(map[string]interface{})

		cloneURLs := links["clone"].([]interface{})

		for _, cloneURL := range cloneURLs {
			curl := cloneURL.(map[string]interface{})

			if curl["name"].(string) == strategy {
				repos = append(repos, curl["href"].(string))
			}
		}
	}

	return repos
}

type bitbucketRemote struct {
	client *bitbucket.Client
	config *config.Bitbucket
}

func (r *bitbucketRemote) ListRepositories() ([]string, error) {
	pageLen := uint64(10)
	allRepos := make([]string, 0)

	strategy := "ssh"
	if r.config.GetStrategy() == config.CloneStrategy_HTTP {
		strategy = "https"
	}

	for _, user := range r.config.Users {
		logrus.Infof("[remotes.bitbucket] fetching projects for user: %s", user)

		for page := uint64(1); true; page++ {
			repos, err := r.client.Repositories.ListForAccount(&bitbucket.RepositoriesOptions{
				ListOptions: &bitbucket.ListOptions{
					Page: page,
					PageLen: pageLen,
				},
				Owner: user,
				Role:  "contributor",
			})

			if err != nil {
				logrus.Errorf("[remotes.bitbucket] encountered err while fetching projects for user %s, %v", user, err)
				continue
			}

			rr := convertRepositoriesResponse(repos, strategy)
			allRepos = append(allRepos, rr...)

			if uint64(len(rr)) < pageLen {
				break
			}
		}
	}

	for _, team := range r.config.Teams {
		logrus.Infof("[remotes.bitbucket] fetching projects for team: %s", team)

		for page := uint64(1); true; page++ {
			repos, err := r.client.Repositories.ListForTeam(&bitbucket.RepositoriesOptions{
				ListOptions: &bitbucket.ListOptions{
					Page: page,
					PageLen: pageLen,
				},
				Owner: team,
				Role:  "contributor",
			})

			if err != nil {
				logrus.Errorf("[remotes.bitbucket] encountered err while fetching projects for team %s, %v", team, err)
				continue
			}

			rr := convertRepositoriesResponse(repos, strategy)
			allRepos = append(allRepos, rr...)

			if uint64(len(rr)) < pageLen {
				break
			}
		}
	}

	return allRepos, nil
}
