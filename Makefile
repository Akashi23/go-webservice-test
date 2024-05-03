#!make
include .env-test
export $(shell sed 's/=.*//' .env-test)
DOCKER_IMAGE_NAME=go-api

build:
	go build -o bin/main .

run: build
	./bin/main

test:
	go test -v ./... -cover

build-docker:
	docker build -t $(DOCKER_IMAGE_NAME) .

run-db:
	docker-compose up -d db
	
up:
	docker-compose up -d

down:
	docker-compose down

clean-compose:
	docker-compose down;
	docker image rm $(DOCKER_IMAGE_NAME)
	rm -rf .dbdata
