# Day 11-15: Advanced Concepts

### Day 11: [Concurrency in Go](/Advanced_Concepts/Day-11/)

Task:

- Learn about goroutines, lightweight threads managed by the Go runtime.
- Understand how to create and manage goroutines for concurrent programming.

Notes:

- Goroutines allow functions to run concurrently with other functions.
- Use the `go` keyword to create a new goroutine.
- Ensure proper synchronization when accessing shared resources.

### Day 12: [Channels and Communication](/Advanced_Concepts/Day-12/)

Task:

- Explore channels, a way to enable communication and synchronization between goroutines.
- Understand the different operations that can be performed on channels.

Notes:

- Channels are a powerful feature for facilitating communication between goroutines.
- They can be used to pass data or synchronize execution.
- Use channel operations such as sending and receiving data and closing channels.

### Day 13: [Defer and Panic](/Advanced_Concepts/Day-13/)

Task:

- Learn about the `defer` statement and how it is used for executing functions after the surrounding function returns.
- Understand the `panic` and `recover` mechanisms for handling exceptional situations.

Notes:

- Use `defer` to ensure cleanup actions or any necessary post-processing.
- `panic` is a built-in function to terminate the program abruptly in exceptional cases.
- `recover` is used to regain control of a panicking goroutine.

### Day 14: [File Handling in Go](/Advanced_Concepts/Day-14/)

Task:

- Explore file handling in Go, including reading from and writing to files.
- Understand file operations and error handling mechanisms related to file handling.

Notes:

- Use the `os` package to perform file-related operations in Go.
- Learn about opening, reading, writing, and closing files.
- Handle errors effectively when performing file operations.

### Day 15: [Testing in Go](/Advanced_Concepts/Day-15/)

Task:

- Learn about testing methodologies in Go and how to write unit tests for Go code.
- Understand the testing package and how to run tests effectively.

Notes:

- Go has a built-in testing package for writing unit tests.
- Write test functions that start with `Test` and take a `*testing.T` argument.
- Run tests using the `go test` command and interpret the results.

