#!/bin/bash
#go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=${pwd}:src/main/proto --micro_out=. --go_out=. src/main/proto/*.proto