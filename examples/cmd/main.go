package main

import (
	"context"
	"log"

	"github.com/Lucascluz/memora-client/client"
)

func main() {
	ctx := context.Background()

	// Create client connection
	memClient, err := client.NewClient("localhost:1212")
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer func() {
		if err := memClient.Close(); err != nil {
			log.Printf("Error closing client: %v", err)
		}
	}()

	testKey := "test"
	testValue := []byte("test value")

	// Set operation
	err = memClient.Set(ctx, testKey, testValue)
	if err != nil {
		log.Printf("Error setting store entry: %v", err)
		return
	}
	log.Printf("✓ Set value to the store")

	// Get operation
	value, err := memClient.Get(ctx, testKey)
	if err != nil {
		log.Printf("Error getting entry: %v", err)
		return
	}
	log.Printf("✓ Got value from store: %s", string(value))

	// Delete operation
	found, err := memClient.Delete(ctx, testKey)
	if err != nil {
		log.Printf("Error deleting entry from the store: %v", err)
		return
	}
	if found {
		log.Printf("✓ Deleted entry from the store")
	} else {
		log.Printf("Entry was not found for deletion")
	}

	// Test string convenience methods
	log.Println("\n--- Testing string convenience methods ---")

	err = memClient.SetString(ctx, "greeting", "Hello, Memora!")
	if err != nil {
		log.Printf("Error setting string: %v", err)
		return
	}
	log.Printf("✓ Set string value")

	greeting, err := memClient.GetString(ctx, "greeting")
	if err != nil {
		log.Printf("Error getting string: %v", err)
		return
	}
	log.Printf("✓ Got string value: %s", greeting)
}
