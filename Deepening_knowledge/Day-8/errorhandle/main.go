package main

import (
	"fmt"
	"errors"
	"os"
)

func sum(a,b int) (result int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("values can't be zero or less than zero")
	}
	return a + b, nil

}

type CustomError struct {
	data string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Occured due to %v", e.data)
}

func main() {
	c, err := sum(10,-1)
	if err != nil {
		var d CustomError
		d.data = "can't assign Zero value "
		fmt.Println(d.Error())
		err := fmt.Errorf("error is : %v", err.Error())
		fmt.Println(err)
		return
	}
	fmt.Println("sum is ",c)

	file,err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(file)

}