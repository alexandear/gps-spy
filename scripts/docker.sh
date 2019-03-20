#!/usr/bin/bash

docker build -t spy-api . -f Dockerfile
docker run --rm -ti spy-api
