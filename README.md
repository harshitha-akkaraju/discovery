# Repository Discovery Service - RDS

Repository Discovery Service (RDS) is a simple gRPC service that encapsulates the logic for discovering repository urls.
This service acts as an aggregator from upstream version control systems (like github, gitlab, bitbucket, etc).
It was named similarly to Envoy's XDS implementation which acts as an aggregate discovery service for services.


## Support

![GitHub](https://img.shields.io/github/license/deps-cloud/rds.svg)
[![Build Status](https://travis-ci.com/deps-cloud/rds.svg?branch=master)](https://travis-ci.com/deps-cloud/rds)
[![Go Report Card](https://goreportcard.com/badge/github.com/deps-cloud/rds)](https://goreportcard.com/report/github.com/deps-cloud/rds)
[![](https://images.microbadger.com/badges/image/depscloud/rds.svg)](https://microbadger.com/images/depscloud/rds)
[![](https://images.microbadger.com/badges/version/depscloud/rds.svg)](https://microbadger.com/images/depscloud/rds)

## Getting Started

Getting started with RDS is really quick and easy.

### Releases

Releases for this project can be found here under github's releases integration.

https://github.com/deps-cloud/rds/releases

### Configuring rds

Once the binary has been installed, you will need to configure `rds`.
By default, `rds` looks for a configuration file at `${HOME}/.rds/config.yml`.
The location of the config file can be changed using the `-config` parameter.
Below, I have provided a snippet on the configuration required for using github as a remote. 

```yaml
# cat ${HOME}/.rds/config.yml
accounts:
- github:
    oauth2:
      token: <oauth_token>
    users:
    - <username>
- gitlab:
    private:
      token: <private_token>
    users:
    - <username>
- bitbucket:
    basic:
      username: <username>
      password: <app_password>
    users:
    - <username>
    teams:
    - <teamname>
```

### Running rds

Once you've set up the configuration, simply run `rds` to start the service.

```bash
$ rds
INFO[0000] [main] starting gRPC on :8090                
INFO[0000] [service.rds] loading repositories           
INFO[0000] [remotes.github] processing organizations for user: mjpitz 
INFO[0000] [remotes.github] processing repositories for user: mjpitz 
INFO[0003] [remotes.github] processing repositories for organization: indeedeng 
INFO[0005] [remotes.github] processing repositories for organization: indeedeng-alpha 
INFO[0005] [remotes.github] processing repositories for organization: deps-cloud 
INFO[0006] [remotes.gitlab] fetching groups
INFO[0006] [remotes.gitlab] fetching projects for user: mjpitz 
INFO[0006] [remotes.gitlab] fetching projects for group: group1
INFO[0007] [remotes.bitbucket] fetching projects for user: mjpitz
INFO[0008] [remotes.bitbucket] fetching projects for team: team1 
```
