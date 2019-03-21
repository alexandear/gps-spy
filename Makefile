default: clean build test

help:
	@echo 'Usage: make <TARGETS> ... <OPTIONS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    build              Compile packages and dependencies.'
	@echo '    build_static       Compile packages and dependencies to static.'
	@echo '    clean              Remove binary.'
	@echo '    dep                Download and install build time dependencies.'
	@echo '    format             Run gofmt on package sources.'
	@echo '    help               Show this help screen.'
	@echo '    test               Run tests.'
	@echo '    docker             Build docker image and run it.'
	@echo '    swagger            Generate server from swagger spec.'
	@echo ''
	@echo 'Targets run by default are: clean build test'
	@echo ''

.PHONY: all clean default build format help test

PKGS = $(shell go list ./... | grep -v /vendor)

BINARY = ./bin/spy-api
IMAGE  = spy-api

all: clean swagger dep format build test docker

build:
	@echo build
	@go build -o $(BINARY)

build_static:
	@echo build static
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $(BINARY) .

clean:
	@echo clean
	@go clean

test:
	@echo test
	@go test -race -count=1 -v ./...

format:
	@echo format
	@go fmt $(PKGS)

dep:
	@echo dep
	@dep ensure -v

docker:
	@echo docker
	@docker build -t $(IMAGE) . -f Dockerfile

docker-run:
	@echo docker run
	@docker run --rm -ti -p 8080:80 $(IMAGE)

swagger:
	@echo swagger
	@rm -rf ./internal/restapi
	@swagger generate server -f ./api/spec.yaml -t ./internal --exclude-main --flag-strategy pflag
