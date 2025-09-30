# Memora - Distributed Cache Store

A simple, fast, and efficient distributed cache store built in Go with gRPC communication.

## 🏗️ Project Structure

This repository contains the complete Memora ecosystem organized as a single repository with multiple modules:

### Core Components

- **[server/](./server/)** - The main cache server implementation
- **[client/](./client/)** - Go client library for easy integration  
- **[proto/](./proto/)** - Protocol Buffer definitions for gRPC communication
- **[examples/](./examples/)** - Example usage demonstrating the client library

## 🚀 Quick Start

### 1. Start the Server

```bash
cd server
go run cmd/main.go
```

The server will start on port `1212` by default.

### 2. Use the Client

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

### 3. Run the Example

```bash
cd examples
go run cmd/main.go
```

## 📚 Documentation

Visit the [complete documentation website](https://lucascluz.github.io/memora) for detailed API reference, examples, and guides.

## 🐳 Docker Deployment

Build and run with Docker:

```bash
# Build the Docker image
docker build -t memora-server .

# Run with Docker
docker run -p 1212:1212 memora-server

# Or use Docker Compose
docker-compose up
```

## 🔧 Development

Each component can be developed independently:

```bash
# Server development
cd server && go run cmd/main.go

# Client development
cd client && go test ./...

# Protocol changes (regenerate from proto directory)
cd proto && protoc --go_out=. --go-grpc_out=. memora.proto
```

## 📦 Module Structure

- **github.com/Lucascluz/memora-server** - Server module
- **github.com/Lucascluz/memora-client** - Client module  
- **github.com/Lucascluz/memora-proto** - Protocol definitions
- **github.com/Lucascluz/memora/examples** - Examples module

## 🎯 Features

- **High Performance**: Built with Go for maximum efficiency
- **Simple API**: Just three operations - Set, Get, Delete
- **gRPC Communication**: Fast, type-safe client-server communication
- **Thread Safe**: Concurrent access protection built-in
- **Easy Integration**: Simple client library with convenience methods
- **Docker Ready**: Easy deployment with Docker and Docker Compose
- **Well Documented**: Complete API reference and examples

## 📄 License

Open source - feel free to use and contribute!

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

---

Built with ❤️ by [Lucas](https://github.com/Lucascluz)

