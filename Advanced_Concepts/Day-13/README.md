# Day 13

## Defer, Panic and Recover

In Go, `panic` and `recover` are mechanisms for handling and managing exceptional situations, primarily dealing with errors and unexpected behavior. Here's a brief explanation of each:

1. **Panic**: When a function encounters an unexpected error or exceptional situation that it cannot handle, it can call the `panic` function. This immediately stops the normal flow of the program, unwinds the stack, and executes deferred functions. A panic is followed by the program's termination and the output of a log message containing the error message and a stack trace. Panics are typically used for unrecoverable errors.

   ```go
   panic("Something went wrong!")
   ```

2. **Recover**: The `recover` function is used to regain control of a panicking goroutine. It is typically used in deferred functions to capture and handle panics. When called from within a deferred function, `recover` stops the propagation of the panic and returns the value passed to the `panic` function. This allows you to handle the panic gracefully and continue the execution of the program.

   ```go
   defer func() {
       if r := recover(); r != nil {
           fmt.Println("Recovered from panic:", r)
       }
   }()
   ```

Using `panic` and `recover` can help you manage unexpected errors and prevent your program from terminating abruptly. However, it's essential to use them judiciously and ensure that your program is designed to handle errors gracefully wherever possible.
