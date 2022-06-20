#!/usr/bin/env bash

go get github.com/googleapis/googleapis

googleapis_ver=latest
googleapis_pkg=github.com/googleapis/googleapis
googleapis_path=$(go list -m -f '{{.Dir}}' ${googleapis_pkg}@${googleapis_ver})

protoc -I . \
    -I ${googleapis_path} \
    --include_imports \
    --go_out=./gen --go_opt=paths=source_relative \
    --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./gen --grpc-gateway_opt paths=source_relative \
    --descriptor_set_out=../../docker/envoy/descriptors/access.pb \
    *.proto



