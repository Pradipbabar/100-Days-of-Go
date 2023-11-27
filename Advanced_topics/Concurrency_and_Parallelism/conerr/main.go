package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type TaskResult struct {
	Result int
	Err    error
}

func simulateTask(id int, resultChan chan<- TaskResult, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate task with a possibility of an error
	if id%2 == 0 {
		// Introduce an error for even IDs
		resultChan <- TaskResult{Result: 0, Err: errors.New("simulated error")}
		return
	}

	// Simulate successful task
	result := id * 2
	resultChan <- TaskResult{Result: result, Err: nil}
}

func main() {
	numTasks := 5
	resultChan := make(chan TaskResult, numTasks)
	var wg sync.WaitGroup

	// Start concurrent tasks
	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go simulateTask(i, resultChan, &wg)
	}

	// Wait for all tasks to finish
	wg.Wait()

	// Close the result channel
	close(resultChan)

	// Collect and handle errors
	for result := range resultChan {
		if result.Err != nil {
			fmt.Printf("Error in task: %v\n", result.Err)
		} else {
			fmt.Printf("Task result: %d\n", result.Result)
		}
	}
}
