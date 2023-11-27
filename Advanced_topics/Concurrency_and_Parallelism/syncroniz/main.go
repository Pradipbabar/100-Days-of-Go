package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter.increment(&wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Printf("Final Counter Value: %d\n", counter.value)
}
