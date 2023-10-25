package main

import (
	"fmt"
)

func like(i interface{}){
	switch i.(type) {
	case string:
		fmt.Printf("the string is %s\n",i.(string))
	case int:
		fmt.Printf("the integer is %d \n",i.(int))
	default:
		fmt.Printf("Unknown Type\n")
	
}
}

func main(){
	like("Hello")
	like(11)
	like(123.45)
}