SHELL := /bin/bash

tidy: # @HELP go mod tidy and update vendor
tidy: 
	go mod tidy
	go mod vendor

test: # @HELP go run application test
test:
	go test ./... -count=1 -v

build: # @HELP build application binary and place in bin directory
build:
	go build \
		-o bin/simple-server \
		./cmd/simple-server

docker: # @HELP build application and run in docker
docker:
	docker build --rm -t app .
	docker-compose up --build