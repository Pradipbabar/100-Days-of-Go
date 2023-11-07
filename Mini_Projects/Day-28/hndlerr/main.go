package main

import (
	"errors"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("invalid-url") // simulate an error
	if err != nil {
		// Wrap the error with additional context
		wrappedErr := fmt.Errorf("error in handler: %w", err)
		fmt.Println(wrappedErr)

		// Chain the error with the original error
		chainErr := errors.New("error in handler: " + err.Error())
		fmt.Println(chainErr)

		// Simulate a panic and recover from it
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		panic("Critical error occurred")
	}
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
