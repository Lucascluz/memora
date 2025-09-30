---
layout: default
title: Memora - Distributed Cache Store
description: A simple, fast, and efficient distributed cache store built in Go
---

# Memora - Distributed Cache Store

A simple, fast, and efficient distributed cache store built in Go with gRPC communication.

## Why Memora?

- **🚀 High Performance** - Built with Go for maximum performance and minimal latency
- **🎯 Simple API** - Clean and intuitive API with just three essential operations
- **🔌 gRPC Powered** - Uses gRPC for fast, type-safe communication
- **📦 Easy Integration** - Simple Go client library that abstracts all complexity
- **🛡️ Thread Safe** - Built-in concurrency control ensures data consistency
- **🌐 Distributed Ready** - Designed to work across multiple nodes

## Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/Lucascluz/memora.git
cd memora

# Add client to your project
go get github.com/Lucascluz/memora-client
```

### Start the Server

```bash
# Run the server from source
cd server
go run cmd/main.go

# Or build and run
go build -o memora-server cmd/main.go
./memora-server

# Or with Docker
docker build -t memora-server .
docker run -p 1212:1212 memora-server
```

### Use the Client

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

### Client Methods

#### `NewClient(address string) (*Client, error)`
Creates a new client connection to the Memora server at the specified address.

```go
client, err := client.NewClient("localhost:1212")
```

#### `Set(ctx context.Context, key string, value []byte) error`
Stores a key-value pair in the cache. If the key already exists, it will be overwritten.

```go
err := client.Set(ctx, "mykey", []byte("myvalue"))
```

#### `Get(ctx context.Context, key string) ([]byte, error)`
Retrieves the value associated with the given key. Returns an error if the key is not found.

```go
value, err := client.Get(ctx, "mykey")
```

#### `Delete(ctx context.Context, key string) (bool, error)`
Removes the key-value pair from the cache. Returns true if the key was found and deleted.

```go
found, err := client.Delete(ctx, "mykey")
```

#### `SetString(ctx context.Context, key, value string) error`
Convenience method to store a string value. Automatically converts to bytes.

```go
err := client.SetString(ctx, "greeting", "Hello World")
```

#### `GetString(ctx context.Context, key string) (string, error)`
Convenience method to retrieve a string value. Automatically converts from bytes.

```go
greeting, err := client.GetString(ctx, "greeting")
```

#### `Close() error`
Closes the connection to the server. Should be called when done using the client.

```go
err := client.Close()
```

## Architecture

The project is organized into several modules:

- **server/** - The core cache server built in Go
- **client/** - Go client library that abstracts gRPC communication  
- **proto/** - Protocol Buffer definitions for gRPC communication
- **examples/** - Example implementations demonstrating usage

## Development

### Building

```bash
# Build all components
make build

# Run tests
make test

# Run integration test
make integration-test
```

### Docker

```bash
# Build Docker image
make docker-build

# Run with Docker Compose
make docker-up
```

## Links

- [GitHub Repository](https://github.com/Lucascluz/memora)
- [Server Module](https://github.com/Lucascluz/memora/tree/main/server)
- [Client Module](https://github.com/Lucascluz/memora/tree/main/client)
- [Protocol Definitions](https://github.com/Lucascluz/memora/tree/main/proto)

---

Built with ❤️ by [Lucas](https://github.com/Lucascluz)