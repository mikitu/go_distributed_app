#!/bin/sh
export GOPATH=$(pwd) && go build  -o ./bin/sensor src/distributed/sensors/sensor.go && ./bin/sensor