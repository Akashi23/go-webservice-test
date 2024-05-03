#!make
include .env-test
export $(shell sed 's/=.*//' .env-test)
DOCKER_IMAGE_NAME=go-api

build:
	go build -o bin/main .

run: build
	./bin/main

build-docker:
	docker build -t $(DOCKER_IMAGE_NAME) .

run-docker-compose:
	docker-compose up -d