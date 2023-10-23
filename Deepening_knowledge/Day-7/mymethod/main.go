package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	Pradip := User{"Pradip", "Pradip@go.dev", true, 16}
	fmt.Println(Pradip)
	fmt.Printf("Pradip details are: %+v\n", Pradip)
	fmt.Printf("Name is %v and email is %v.\n", Pradip.Name, Pradip.Email)
	Pradip.GetStatus()
	Pradip.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", Pradip.Name, Pradip.Email)

	    // Initializing the values
    // of the author structure
    res := author{
        name:   "Sona",
        branch: "CSE",
    }
 
    fmt.Println("Branch Name(Before): ", res.branch)
 
    // Calling the show_1 method
    // (pointer method) with value
    res.show_1("ECE")
    fmt.Println("Branch Name(After): ", res.branch)
 
    // Calling the show_2 method
    // (value method) with a pointer
    (&res).show_2()
    fmt.Println("Author's name(After): ", res.name)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Author structure
type author struct {
    name   string
    branch string
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}
 
// Method with a value
// receiver of author type
func (a author) show_2() {
 
    a.name = "Gourav"
    fmt.Println("Author's name(Before) : ", a.name)
}