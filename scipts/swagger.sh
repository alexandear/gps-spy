#!/usr/bin/bash

swagger generate server -f ./api/spec.yaml -t ./internal --exclude-main
