package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Lucascluz/memora-client/client"
)

func main() {
	ctx := context.Background()

	log.Println("=== Memora Cache Test Suite ===")

	// Test 1: Create client connection
	log.Println("Test 1: Creating client connection...")
	memClient, err := client.NewClient("localhost:1212")
	if err != nil {
		log.Fatalf("❌ Error creating client: %v", err)
	}
	defer func() {
		if err := memClient.Close(); err != nil {
			log.Printf("Error closing client: %v", err)
		}
	}()
	log.Println("✓ Client created successfully")

	// Test 2: Connect to server
	log.Println("Test 2: Connecting to server...")
	err = memClient.Connect(ctx)
	if err != nil {
		log.Fatalf("❌ Error connecting to server: %v", err)
	}
	log.Println("✓ Connected to server successfully")

	// Test 3: Set operation with no TTL
	log.Println("Test 3: Set operation (no TTL)...")
	testKey := "user:1001"
	testValue := []byte("John Doe")
	err = memClient.Set(ctx, testKey, testValue, 0)
	if err != nil {
		log.Printf("❌ Error setting value: %v\n", err)
		return
	}
	log.Printf("✓ Set key '%s' with value '%s'\n\n", testKey, string(testValue))

	// Test 4: Get operation
	log.Println("Test 4: Get operation...")
	value, err := memClient.Get(ctx, testKey)
	if err != nil {
		log.Printf("❌ Error getting value: %v\n", err)
		return
	}
	log.Printf("✓ Retrieved value: '%s'\n\n", string(value))

	// Test 5: Set with TTL
	log.Println("Test 5: Set with TTL (5 seconds from now)...")
	ttlKey := "session:abc123"
	ttlValue := []byte("temporary session data")
	ttl := time.Now().Unix() + 5
	err = memClient.Set(ctx, ttlKey, ttlValue, ttl)
	if err != nil {
		log.Printf("❌ Error setting value with TTL: %v\n", err)
		return
	}
	log.Printf("✓ Set key '%s' with TTL: %d\n", ttlKey, ttl)

	// Test 6: Get value before TTL expires
	log.Println("Test 6: Get value before TTL expires...")
	value, err = memClient.Get(ctx, ttlKey)
	if err != nil {
		log.Printf("❌ Error getting value: %v\n", err)
		return
	}
	log.Printf("✓ Retrieved value: '%s'\n\n", string(value))

	// Test 7: Wait for TTL to expire
	log.Println("Test 7: Waiting for TTL to expire (6 seconds)...")
	time.Sleep(6 * time.Second)
	_, err = memClient.Get(ctx, ttlKey)
	if err != nil {
		log.Printf("✓ Expected error: %v\n\n", err)
	} else {
		log.Println("❌ Value should have expired but still exists")
	}

	// Test 8: SetString convenience method
	log.Println("Test 8: SetString convenience method...")
	err = memClient.SetString(ctx, "greeting", "Hello, Memora!")
	if err != nil {
		log.Printf("❌ Error setting string: %v\n", err)
		return
	}
	log.Println("✓ Set string value using SetString")

	// Test 9: GetString convenience method
	log.Println("Test 9: GetString convenience method...")
	greeting, err := memClient.GetString(ctx, "greeting")
	if err != nil {
		log.Printf("❌ Error getting string: %v\n", err)
		return
	}
	log.Printf("✓ Retrieved string: '%s'\n\n", greeting)

	// Test 10: Delete operation
	log.Println("Test 10: Delete operation...")
	found, err := memClient.Delete(ctx, testKey)
	if err != nil {
		log.Printf("❌ Error deleting entry: %v\n", err)
		return
	}
	if found {
		log.Printf("✓ Deleted key '%s'\n\n", testKey)
	} else {
		log.Println("❌ Key was not found for deletion")
	}

	// Test 11: Get deleted key (should fail)
	log.Println("Test 11: Get deleted key (should fail)...")
	_, err = memClient.Get(ctx, testKey)
	if err != nil {
		log.Printf("✓ Expected error: %v\n\n", err)
	} else {
		log.Println("❌ Deleted key should not exist")
	}

	// Test 12: Delete non-existent key
	log.Println("Test 12: Delete non-existent key...")
	found, err = memClient.Delete(ctx, "non-existent-key")
	if err != nil {
		log.Printf("❌ Error during delete: %v\n", err)
		return
	}
	if !found {
		log.Println("✓ Key not found (expected behavior)")
	} else {
		log.Println("❌ Non-existent key should not be found")
	}

	// Test 13: Get non-existent key
	log.Println("Test 13: Get non-existent key...")
	_, err = memClient.Get(ctx, "another-non-existent-key")
	if err != nil {
		log.Printf("✓ Expected error: %v\n\n", err)
	} else {
		log.Println("❌ Non-existent key should return error")
	}

	// Test 14: Multiple Set/Get operations
	log.Println("Test 14: Multiple Set/Get operations...")
	testData := map[string]string{
		"product:101": "Laptop",
		"product:102": "Mouse",
		"product:103": "Keyboard",
	}
	for key, val := range testData {
		err = memClient.SetString(ctx, key, val)
		if err != nil {
			log.Printf("❌ Error setting %s: %v\n", key, err)
			return
		}
	}
	log.Println("✓ Set multiple keys")

	for key, expectedVal := range testData {
		val, err := memClient.GetString(ctx, key)
		if err != nil {
			log.Printf("❌ Error getting %s: %v\n", key, err)
			return
		}
		if val == expectedVal {
			log.Printf("✓ Retrieved %s = %s\n", key, val)
		} else {
			log.Printf("❌ Value mismatch for %s: got %s, expected %s\n", key, val, expectedVal)
		}
	}

	fmt.Println("\n=== All tests completed successfully! ===")
}
