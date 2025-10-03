// Package client provides a simple Go client for the Memora distributed cache service.
// It abstracts the gRPC communication and provides easy-to-use methods for cache operations.
package client

import (
	"context"
	"fmt"
	"net"

	pb "github.com/Lucascluz/memora-proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.MemoraServiceClient
	key    string
}

// NewClient creates a new gRPC client connection to the Memora service at the specified address.
// It establishes an insecure connection and returns a Client instance or an error if connection fails.
func NewClient(address string) (*Client, error) {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server at %s: %w", address, err)
	}

	// return connection to the grpc server
	c := pb.NewMemoraServiceClient(conn)

	return &Client{conn: conn, client: c, key: ""}, nil
}

// Connect establishes a connection with the Memora server and gets a client key
func (c *Client) Connect(ctx context.Context) error {
	// Get the local IP address
	clientIP, err := getLocalIP()
	if err != nil {
		return fmt.Errorf("failed to get local IP: %w", err)
	}

	req := &pb.ConnectionRequest{ClientIP: clientIP}

	resp, err := c.client.Connect(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}

	if !resp.Success {
		return fmt.Errorf("connection failed")
	}

	// Store the client key for future requests
	c.key = resp.ClientKey
	return nil
}

// Set stores a key-value pair in the Memora service.
// It takes a context, key string, and value as bytes, returning an error if the operation fails.
func (c *Client) Set(ctx context.Context, key string, value []byte, ttl int64) error {
	req := &pb.SetRequest{ClientKey: c.key, EntryKey: key, Value: value, Ttl: ttl}
	resp, err := c.client.Set(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to set key %s: %w", key, err)
	}
	if !resp.Success {
		return fmt.Errorf("set operation failed for key %s: %s", key, resp.Status)
	}
	return nil
}

// Get retrieves the value associated with the given key from the Memora service.
// It returns the value as bytes if found, or an error if the key doesn't exist or operation fails.
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	req := &pb.GetRequest{ClientKey: c.key, EntryKey: key}
	resp, err := c.client.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get key %s: %w", key, err)
	}
	if resp.Status != "found" {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return resp.Value, nil
}

// Delete removes the key-value pair from the Memora service.
// It returns true if the key was found and deleted, false otherwise, along with any error.
func (c *Client) Delete(ctx context.Context, key string) (bool, error) {
	req := &pb.DeleteRequest{ClientKey: c.key, EntryKey: key}
	resp, err := c.client.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("failed to delete key %s: %w", key, err)
	}
	return resp.Found, nil
}

// Close terminates the gRPC connection to the Memora service.
// It returns an error if the connection fails to close properly.
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// SetString stores a string value in the cache with the given key.
// This is a convenience method that converts the string to bytes.
// The ttl is set to 0 (no expiration) by default.
func (c *Client) SetString(ctx context.Context, key, value string) error {
	return c.Set(ctx, key, []byte(value), 0)
}

// GetString retrieves a value from the cache and returns it as a string.
// This is a convenience method that converts bytes to string.
func (c *Client) GetString(ctx context.Context, key string) (string, error) {
	data, err := c.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// getLocalIP returns the local IP address of the client
func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
