BASEPATH = $(shell pwd)
export PATH := $(BASEPATH)/bin:$(PATH)

.PHONY: all clean default format help test

default: clean build test

help:
	@echo 'Usage: make <TARGETS> ... <OPTIONS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    build              Compile packages and dependencies.'
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

PKGS        = $(shell go list ./... | grep -v /vendor)
SCRIPTS_DIR = ./scripts

all: clean swagger dep format build test docker

build:
	@echo build
	@. $(SCRIPTS_DIR)/build.sh

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
	@. $(SCRIPTS_DIR)/docker.sh

swagger:
	@echo swagger
	@. $(SCRIPTS_DIR)/swagger.sh
