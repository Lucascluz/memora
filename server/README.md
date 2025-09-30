# Memora Server

The core cache server component of the Memora distributed cache system.

## Features

- **High Performance**: Built with Go for optimal speed and efficiency
- **gRPC API**: Fast, type-safe communication protocol
- **Thread Safe**: Concurrent access protection with mutex locks
- **Simple Operations**: Set, Get, Delete operations
- **Memory Efficient**: In-memory storage with minimal overhead

## Installation

```bash
go install github.com/Lucascluz/memora-server@latest
```

Or run from source:

```bash
go run cmd/main.go
```

## Configuration

The server runs on port `1212` by default. You can modify the port in `cmd/main.go`.

## API

The server implements the following gRPC methods:

- `Set(SetRequest) returns (SetResponse)` - Store a key-value pair
- `Get(GetRequest) returns (GetResponse)` - Retrieve a value by key
- `Delete(DeleteRequest) returns (DeleteResponse)` - Remove a key-value pair

## Development

```bash
# Install dependencies
go mod tidy

# Run the server
go run cmd/main.go

# Run tests
go test ./...
```

## Architecture

```
cmd/
├── main.go              # Server entry point
internal/
├── cache/
│   └── cache.go         # Cache implementation
└── server/
    └── server.go        # gRPC server implementation
```

The server uses a simple in-memory map protected by mutex for thread-safe operations.

## Performance

The server is designed to handle thousands of concurrent requests efficiently with minimal latency.