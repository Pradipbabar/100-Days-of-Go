package main

import (
	"fmt"
)
func main() {
	fmt.Println("Hello")
	greet()

	result := sum(7,9)
	fmt.Println("the sum is : ", result) 

	proRes, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)

}

func greet() {
	fmt.Println("Hi ")
}

func sum(valone int, valtwo int) int {

	return valone + valtwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"

}