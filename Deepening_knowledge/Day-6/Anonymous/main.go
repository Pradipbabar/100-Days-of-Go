// Go program to pass an anonymous 
// function as an argument into 
// other function 
package main 

import "fmt"


// Passing anonymous function 
// as an argument 
func GFG(i func(p, q string)string){ 
	fmt.Println(i ("Learn_", "for")) 
	
} 
	
func main() { 
	value:= func(p, q string) string{ 
		return p + q + "_Earn"
	} 
	GFG(value) 
} 
