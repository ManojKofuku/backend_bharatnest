#!/bin/bash

# Exit on error
set -e

echo "Installing Protocol Buffer compiler plugins for Go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Make sure Go bin is in PATH
export PATH=$PATH:$(go env GOPATH)/bin

echo "Generating Go code from Protocol Buffers..."
protoc --proto_path=backend/proto \
       --go_out=backend/proto --go_opt=module=backend/proto \
       --go-grpc_out=backend/proto --go-grpc_opt=module=backend/proto \
       backend/proto/hotel.proto

echo "Protocol Buffers code generation complete!" 