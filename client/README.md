# Memora Client

Go client library for the Memora distributed cache system. This client provides a simple, efficient way to interact with Memora cache servers using gRPC.

## Installation

```bash
go get github.com/Lucascluz/memora-client
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    
    "github.com/Lucascluz/memora-client/client"
)

func main() {
    // Create client
    memClient, err := client.NewClient("localhost:1212")
    if err != nil {
        log.Fatal(err)
    }
    defer memClient.Close()
    
    ctx := context.Background()
    
    // Set a value
    err = memClient.Set(ctx, "key1", []byte("Hello, Memora!"))
    if err != nil {
        log.Fatal(err)
    }
    
    // Get the value
    value, err := memClient.Get(ctx, "key1")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Value: %s", string(value))
    
    // Delete the key
    found, err := memClient.Delete(ctx, "key1")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Deleted: %v", found)
}
```

## API Reference

### Core Methods

- **`NewClient(address string) (*Client, error)`** - Create new client connection
- **`Set(ctx context.Context, key string, value []byte) error`** - Store key-value pair
- **`Get(ctx context.Context, key string) ([]byte, error)`** - Retrieve value by key
- **`Delete(ctx context.Context, key string) (bool, error)`** - Remove key-value pair
- **`Close() error`** - Close the connection

### Convenience Methods

- **`SetString(ctx context.Context, key, value string) error`** - Store string value
- **`GetString(ctx context.Context, key string) (string, error)`** - Retrieve string value

## Project Structure

```
memora-client/
├── client/
│   └── client.go       # Client implementation
├── examples/
│   └── main.go         # Example usage
├── go.mod              # Go module configuration
├── go.sum              # Dependency checksums
└── README.md           # Documentation
```

## Error Handling

All methods return descriptive errors with context:

```go
value, err := client.Get(ctx, "nonexistent")
if err != nil {
    log.Printf("Get failed: %v", err) // "Get failed: key nonexistent not found"
}
```

## Connection Management

The client maintains a persistent gRPC connection. Always call `Close()` when done:

```go
defer func() {
    if err := client.Close(); err != nil {
        log.Printf("Error closing client: %v", err)
    }
}()
```

## Features

- **Thread Safe**: Can be used concurrently from multiple goroutines
- **Error Context**: Descriptive error messages with operation context
- **Convenience Methods**: String helpers for common use cases
- **Resource Management**: Proper connection lifecycle management
- **Type Safety**: Leverages gRPC for type-safe communication

## Examples

See the [examples directory](./examples/) for complete usage examples.