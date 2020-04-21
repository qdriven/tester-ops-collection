#!/usr/bin/env sh

MODULE=$1

cd {$MODULE}
CGO_ENABLED=0
## TODO build in MAC
GOOS=linux GOARCH=amd64 vgo build -ldflags "-s -w"
GOOS=windows GOARCH=amd64 vgo build -ldflags "-s -w"