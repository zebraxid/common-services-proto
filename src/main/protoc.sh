#!/bin/bash
#go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=${pwd}:proto --micro_out=go --go_out=go proto/*.proto