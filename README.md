# Repository Discovery Service - RDS

[![Build Status](https://travis-ci.com/mjpitz/rds.svg?branch=master)](https://travis-ci.com/mjpitz/rds)
[![Go Report Card](https://goreportcard.com/badge/github.com/mjpitz/rds)](https://goreportcard.com/report/github.com/mjpitz/rds)

Repository Discovery Service (RDS) is a simple gRPC service that encapsulates the logic for discovering repository urls.
This service acts as an aggregator from upstream version control systems (like github, gitlab, bitbucket, etc).
It was named similarly to Envoy's XDS implementation which acts as an aggregate discovery service for services.
 