#!/bin/bash
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. *.proto