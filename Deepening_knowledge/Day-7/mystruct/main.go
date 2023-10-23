package main

import "fmt"

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
  }

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	Pradip := User{"Pradip", "Pradip@go.dev", true, 22}
	fmt.Println(Pradip)
	fmt.Printf("Pradip details are: %+v\n", Pradip)
	fmt.Printf("Name is %v and email is %v.", Pradip.Name, Pradip.Email)

	var pers1 Person
	var pers2 Person
  
	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
  
	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
  
	// Print Pers1 info by calling a function
	printPerson(pers1)
  
	// Print Pers2 info by calling a function
	printPerson(pers2)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Person struct {
	name string
	age int
	job string
	salary int
  }