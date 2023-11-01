package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"net"
  
)

const url = "https://github.com/Pradipbabar/100-Days-of-Go"


func main()  {
	fmt.Println("Hello")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Response",response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	containt := string(databytes)
	// fmt.Println(containt)

	http.HandleFunc("/", func (w http.ResponseWriter, 
        r *http.Request) {
  
        // when receive the request, print the greeting meassage
        fmt.Fprint(w, containt)
  
      })

  // print out the server is going to start and show the time
  log.Println("Starting server....")

  // create server at localhost:8080 and using tcp as the network
  listener, err := net.Listen("tcp", ":8080")

  // if recieve error, record it and exit the program
  if err != nil {
    log.Fatal(err)
  }

  // setup HTTP connection for the listener of the server
  http.Serve(listener, nil)


}