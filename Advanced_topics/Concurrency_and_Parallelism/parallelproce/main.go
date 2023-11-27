package main

import (
	"fmt"
	"sync"
)

func calculateSquare(number int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	result <- square
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, resultChan, &wg)
	}

	// Close the result channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Receive and print results from the channel
	for square := range resultChan {
		fmt.Printf("Square: %d\n", square)
	}
}
