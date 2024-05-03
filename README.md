# go-webservice-test

[![Go-App-CI/CD](https://github.com/Akashi23/go-webservice-test/actions/workflows/go.yaml/badge.svg)](https://github.com/Akashi23/go-webservice-test/actions/workflows/go.yaml)

## Description
This is a simple web service for checking IIN (Individual Identification Number) validity and creating, getting with IINs.

## Installation
1. Clone the repository
2. Run `go mod download` to download dependencies
3. .env-test file for testing and you may use it for testing
4. Run `make run` to start the server
5. Run `make test` to run tests

## Make commands
- `make run` - run the server
- `make test` - run tests
- `make build` - build the project
- `make build-docker` - build the docker image
- `make up` - run the docker container
- `make down` - stop the docker container
- `make clean-compose` - remove the docker container

