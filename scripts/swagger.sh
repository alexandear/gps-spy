#!/usr/bin/bash

rm -rf ./internal/restapi
swagger generate server -f ./api/spec.yaml -t ./internal --exclude-main
