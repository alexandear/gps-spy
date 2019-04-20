MAKEFILE_PATH := $(abspath $(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
PATH := $(MAKEFILE_PATH):$(PATH)

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
SPEC   = ./api/spec.yaml
GEN    = ./internal/gen

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
	@CGO_ENABLED=0 go test -count=1 -v ./...

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
	@rm -rf $(GEN)/models
	@rm -rf $(GEN)/restapi
	@swagger generate server -f $(SPEC) -t $(GEN) --exclude-main --flag-strategy pflag

doc:
	@echo swagger doc
	@swagger serve --flavor=swagger $(SPEC)

generate:
	@echo mock
ifeq ("$(wildcard ./bin/mockery)","")
	@echo build mockery
	@go build -o ./bin/mockery ./vendor/github.com/vektra/mockery/cmd/mockery/
endif
	@go generate ./internal/service/... ./internal/storage/...

grpc:
ifeq ("$(wildcard protoc-gen-go)","")
	@go build -o protoc-gen-go ./vendor/github.com/golang/protobuf/protoc-gen-go
endif
	@protoc \
		-I ./internal/service/specnomery \
		./internal/service/specnomery/grpc.proto \
		--go_out=plugins=grpc:./internal/service/specnomery
