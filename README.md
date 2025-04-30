# Hotel API with gRPC

This project provides both REST API (using Gin) and gRPC endpoints for hotel management.

## Setup

1. Install the Protocol Buffers compiler:

   ```
   # macOS
   brew install protobuf

   # Linux
   apt install -y protobuf-compiler

   # Windows (using Chocolatey)
   choco install protoc
   ```

2. Install Go Protocol Buffers plugins:

   ```
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

   Make sure your PATH includes `$(go env GOPATH)/bin`

3. Generate the protobuf Go code:

   ```
   ./backend/generate.sh
   ```

4. Install the required Go dependencies:

   ```
   go mod tidy
   ```

5. Run the server:
   ```
   go run backend/main.go
   ```

## Available Services

### gRPC Services (port 50051)

- **GetAllHotels**: Retrieves a list of all hotels with pagination support
  - Request: `GetAllHotelsRequest` with optional pagination parameters
  - Response: `GetAllHotelsResponse` with hotel details and total count

### REST API (port 8080)

- The existing REST endpoints continue to work as before

## Example Client Usage

An example client is provided in `backend/client/client.go`. Run it with:

```
go run backend/client/client.go
```

## Project Structure

- `backend/proto/hotel.proto`: Protocol Buffers definition file
- `backend/proto/hotel/*.pb.go`: Generated Go code from the proto file
- `backend/controllers/hotelController.go`: gRPC service implementation
- `backend/client/client.go`: Example client for testing the gRPC service

## Data Mapping

The service automatically maps between the GORM models and Protocol Buffer messages, including:

- Converting database JSON (amenities) to/from protocol buffer repeated fields
- Handling pagination
- Converting timestamp values

## Troubleshooting

If you encounter import errors, make sure:

1. You've generated the protobuf code with `./backend/generate.sh`
2. The import paths in your code match the paths generated in the protocol buffer files
3. You've run `go mod tidy` to update dependencies
