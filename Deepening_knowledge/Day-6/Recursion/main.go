package main

import "fmt"

// Factorial Using Recursion in Go
func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}

// Return Fibonacci Series 
func fibonaci(i int) (ret int) {
	if i == 0 {
	   return 0
	}
	if i == 1 {
	   return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
 }

func main() { 
   var i int = 15
   fmt.Printf("Factorial of %d is %d", i, factorial(i))

   var j int
   for j = 0; j < 10; j++ {
      fmt.Printf("  %d ", fibonaci(j))
   } 
}