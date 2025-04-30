#!/bin/bash

# Exit on error
set -e

echo "Generating Protocol Buffer code for Hotel API..."

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo "Error: protoc is not installed. Please install Protocol Buffers compiler."
    echo "Visit: https://grpc.io/docs/protoc-installation/"
    exit 1
fi

# Check if Go plugins are installed
if ! protoc --version 2>&1 | grep -q "libprotoc"; then
    echo "Error: Protocol Buffers compiler (protoc) is required."
    exit 1
fi

# Create target directories if they don't exist
mkdir -p backend/pb/hotel

# Generate Go code from proto
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       backend/proto/hotel.proto

echo "Code generation complete!"
echo "Generated files in ./backend/pb/hotel/"

# Make the script executable
chmod +x ./scripts/gen_proto.sh 