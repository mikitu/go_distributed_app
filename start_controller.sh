#!/bin/sh
export GOPATH=$(pwd) && go build -o ./bin/controller src/distributed/controller/exec/main.go  && ./bin/controller "$@"