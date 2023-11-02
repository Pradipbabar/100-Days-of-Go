package main

import (
	"fmt"
	 "net/http"
 	"html/template"
 	"strconv"
)

type Student struct {
	Id   int
	Name string //exported field since it begins with a capital letter
}

var id = 1

func indexHTMLTemplateVariableHandler(response http.ResponseWriter, request *http.Request) {
	tmplt,  err := template.ParseFiles("templates/index.html")
	if err != nil {
	  fmt.Println(err) // Ugly debug output
	  return
	}

	if request.Method == "POST" {
		request.ParseForm()
		id := request.FormValue("id")
		name := request.FormValue("name")

		studentId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(response, "Invalid Id", http.StatusBadRequest)
			return
		} 

	p := Student{Id: studentId, Name: name} //define an instance with required field

	tmplt.Execute(response, p) //merge template ‘t’ with content of ‘p’

	}
}


func main() {
	fmt.Println("Starting Server for Templated response from file")

	http.HandleFunc("/", indexHTMLTemplateVariableHandler)

	http.ListenAndServe(":8080", nil)
}