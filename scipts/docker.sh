#!/usr/bin/bash

docker build -t spy-api:latest . -f Dockerfile
docker run --rm -ti spy-api
