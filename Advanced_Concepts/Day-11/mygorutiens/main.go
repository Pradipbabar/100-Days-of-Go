package main

import (
    "fmt"
    "time"

)

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}


func main() {

	go greeter("Hello")
	greeter("world")

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")



}