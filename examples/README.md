# Memora Example

A simple example demonstrating how to use the Memora client library to interact with a Memora cache server.

## Overview

This example shows the basic operations of the Memora cache system:
- Setting key-value pairs
- Retrieving values by key
- Deleting cache entries
- Using string convenience methods

## Prerequisites

1. **Start the Memora Server**:
   ```bash
   cd ../memora-server
   go run cmd/main.go
   ```
   The server should be running on `localhost:1212`

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

## Running the Example

```bash
go run cmd/main.go
```

## Expected Output

```
2025/09/30 17:19:21 ✓ Set value to the store
2025/09/30 17:19:21 ✓ Got value from store: test value
2025/09/30 17:19:21 ✓ Deleted entry from the store

--- Testing string convenience methods ---
2025/09/30 17:19:21 ✓ Set string value
2025/09/30 17:19:21 ✓ Got string value: Hello, Memora!
```

## Code Walkthrough

The example demonstrates:

1. **Client Creation**: Establishing connection to the server
   ```go
   memClient, err := client.NewClient("localhost:1212")
   ```

2. **Setting Data**: Storing a byte array value
   ```go
   err = memClient.Set(ctx, "test", []byte("test value"))
   ```

3. **Getting Data**: Retrieving the stored value
   ```go
   value, err := memClient.Get(ctx, "test")
   ```

4. **Deleting Data**: Removing an entry from cache
   ```go
   found, err := memClient.Delete(ctx, "test")
   ```

5. **String Convenience**: Using string helper methods
   ```go
   err = memClient.SetString(ctx, "greeting", "Hello, Memora!")
   greeting, err := memClient.GetString(ctx, "greeting")
   ```

6. **Proper Cleanup**: Closing the client connection
   ```go
   defer memClient.Close()
   ```

## Error Handling

The example includes proper error handling for each operation:

```go
if err != nil {
    log.Printf("Error: %v", err)
    return
}
```

## Customization

You can modify this example to:
- Connect to a different server address
- Test different data types
- Implement your own error handling strategies
- Add performance benchmarks
- Test concurrent operations

## Next Steps

After running this example, you can:
1. Integrate the client into your own applications
2. Build more complex caching strategies
3. Implement distributed cache scenarios
4. Add monitoring and metrics