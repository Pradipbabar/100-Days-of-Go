package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "https://raw.githubusercontent.com/Pradipbabar/100-Days-of-Go/main/Advanced_Features/README.md"

func main()  {
	fmt.Println("Hello")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response",response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	containt := string(databytes)
	fmt.Println(containt)
}