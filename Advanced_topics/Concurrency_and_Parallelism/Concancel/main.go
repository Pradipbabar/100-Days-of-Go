package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func longRunningOperation(ctx context.Context, result chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled.")
			return
		default:
			// Simulate some work
			time.Sleep(time.Second)
			// Send a random result to the channel
			result <- rand.Intn(100)
		}
	}
}

func main() {
	// Create a context with cancellation capability
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to receive results from the long-running operation
	resultChan := make(chan int)

	// Start the long-running operation
	go longRunningOperation(ctx, resultChan)

	// Simulate the main program doing some work
	for i := 0; i < 5; i++ {
		result := <-resultChan
		fmt.Printf("Received result: %d\n", result)
	}

	// Cancel the long-running operation after receiving 5 results
	cancel()

	// Wait for a moment to observe the cancellation
	time.Sleep(2 * time.Second)
}
