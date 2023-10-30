package main

import (
	"fmt"
)

func multipychnn (ch chan int) {
	fmt.Println(100 * <-ch)
}

func initstring(chnl chan string) {
	for v := 0; v < 3; v++ {
		chnl <- "Pradip"
	}
	close(chnl)
}

func main () {
	fmt.Println("Hello")
	firstchnn := make(chan int)
	fmt.Println("value of chann is : ",firstchnn)
	fmt.Printf("Type of channel is %T \n",firstchnn)
	
	ch := make(chan int)
	fmt.Println("main function")
	go multipychnn(ch)
	ch <- 10
	fmt.Println("exit main function")

	c := make(chan string)
	go initstring(c)
	for {
		resp, ok := <-c
		if ok == false {
			fmt.Println("channel close", ok)
			break
		}
		fmt.Println("channel Open",resp,ok)
	}

	mychnl := make(chan string, 4)
	mychnl <- "Channel 1"
	mychnl <- "Channel 2"
	mychnl <- "Channel 3"
	mychnl <- "Channel 4"
	fmt.Println("Length of channel is : ",cap(mychnl))

	testch := make(chan string)
	// anonymous go function
	go func() {
		testch <- "value 1"
		testch <- "value 2"
		close(testch)
	}()
	for res := range testch{
		fmt.Println(res)
	}

}

