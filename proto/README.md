# Memora Protocol Definitions

Protocol Buffer definitions for the Memora distributed cache system. This package defines the gRPC service interface and message types used for client-server communication.

## Overview

This package provides the Protocol Buffer (protobuf) definitions that enable type-safe, efficient communication between Memora clients and servers using gRPC.

## Structure

```
memora-proto/
├── proto/
│   ├── memora.proto           # Protocol definitions
│   └── gen/
│       ├── memora.pb.go       # Generated Go types
│       └── memora_grpc.pb.go  # Generated gRPC service
├── go.mod
├── go.sum
└── README.md
```

## Service Definition

The `MemoraService` provides three simple operations:

```protobuf
service MemoraService {
    rpc Set (SetRequest) returns (SetResponse);
    rpc Get (GetRequest) returns (GetResponse);
    rpc Delete (DeleteRequest) returns (DeleteResponse);
}
```

## Message Types

### Set Operation
```protobuf
message SetRequest {
    string key = 1;
    bytes value = 2;
    int64 ttl = 3;  // Reserved for future TTL support
}

message SetResponse {
    bool success = 1;
    string status = 2;
}
```

### Get Operation
```protobuf
message GetRequest {
    string key = 1; 
}

message GetResponse {
    string status = 1;
    bytes value = 2;
}
```

### Delete Operation
```protobuf
message DeleteRequest {
    string key = 1;
}

message DeleteResponse {
    bool found = 1;
    string status = 2;
}
```

## Regenerating Code

If you modify the `.proto` file, regenerate the Go code:

```bash
# Install required tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code
protoc --go_out=. --go-grpc_out=. proto/memora.proto
```

## Usage

Import the generated packages in your Go code:

```go
import pb "github.com/Lucascluz/memora-proto/proto/gen"
```

## Design Principles

- **Simplicity**: Minimal message structure for maximum efficiency
- **Extensibility**: Fields reserved for future enhancements (like TTL)
- **Type Safety**: Protocol Buffers ensure type safety across languages
- **Performance**: Binary serialization for optimal network efficiency

## Dependencies

- Protocol Buffers compiler (`protoc`)
- Go protobuf plugin (`protoc-gen-go`)
- Go gRPC plugin (`protoc-gen-go-grpc`)