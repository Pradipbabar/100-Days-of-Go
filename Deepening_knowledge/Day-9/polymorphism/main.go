package main

import (
	"fmt"
)

type values struct {
	first int
	second int
}

type mathtest interface {
	add(a,b int) int
	multipy(a,b int) int 

}

func (v values) add(a,b int) int  {
	return a + b + v.first + v.second
	
}

func (v values) multipy(a,b int) int {
	return a * b * v.first * v.second
}

func main() {
	var te mathtest = values{1,2}
	fmt.Println(te.add(1,2))
	fmt.Println(te.multipy(1,2))

	var test interface{}
	test = "Some string"
	val, ok := test.(string)
	if ok {
		fmt.Println(val)
	}

}