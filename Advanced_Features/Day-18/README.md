# Day 18

## Simple Web Servers with Go

Building a simple web server using the standard `net/http` package in Go is straightforward. Here's an overview of how to handle HTTP requests and responses to create a basic web server:

1. **Import the `net/http` Package**:

   Import the `net/http` package, which provides the necessary functionality for creating an HTTP server and handling requests.

2. **Define a Handler Function**:

   Define a handler function that processes incoming HTTP requests. This function should implement the `http.Handler` interface and handle the logic for different routes and requests.

3. **Register Routes**:

   Use the `http.HandleFunc` or `http.Handle` function to register routes and associate them with specific handler functions.

4. **Start the Server**:

   Start the server using the `http.ListenAndServe` function, specifying the server address and port to listen on.

Here's an example of a simple web server that responds with "Hello, World!" to all incoming requests:

```go
package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorld)

	fmt.Println("Server is starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

In this example, the `helloWorld` function is registered as the handler for the root route. When the server receives an incoming HTTP request, it calls the `helloWorld` function to handle the request and respond with "Hello, World!". The `http.ListenAndServe` function starts the server and listens for incoming connections on port 8080.
