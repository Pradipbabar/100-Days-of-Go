package main

import (
	"fmt"
)

func main() {
	x := 0
	y := 14
	printalloprations(x,y) 
}

func printalloprations(x int, y int) {
	defer func() {
		if r := recover() ; r != nil {
			fmt.Println("Recovering from panic",r)
			fmt.Println("proceeding")
			skipdivide(x,y)
		}
	}()
	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("sum=%v, divide=%v, multiply=%v \n",sum, divide, multiply)
}

func skipdivide(x int, y int) {
	sum, multiply := x+y, x*y
	fmt.Printf("sum=%v, multiply=%v \n",sum, multiply)
}