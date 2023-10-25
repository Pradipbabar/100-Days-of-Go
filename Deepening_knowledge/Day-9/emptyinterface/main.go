package main

import (
	"fmt"
)

var i interface{}

func found (i interface{}){
	fmt.Printf("Type = %T, Value = %v\n",i,i)
}

func main(){
	s := "Data"
	found(s)
	i := 07
	found(i)
}