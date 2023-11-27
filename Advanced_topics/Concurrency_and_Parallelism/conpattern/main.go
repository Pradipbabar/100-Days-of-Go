package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numWorkers = 3

func generator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(100)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
	close(ch)
}

func worker(id int, input <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range input {
		fmt.Printf("Worker %d processing: %d\n", id, num)
		// Simulate some processing
		time.Sleep(time.Millisecond * 1000)
		resultChan <- num * 2 // Double the number
	}
}

func main() {
	inputChan := make(chan int, numWorkers)
	resultChan := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// Start generator
	wg.Add(1)
	go generator(inputChan, &wg)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, inputChan, resultChan, &wg)
	}

	// Start single worker for result processing
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Print results
	for result := range resultChan {
		fmt.Printf("Processed result: %d\n", result)
	}
}
