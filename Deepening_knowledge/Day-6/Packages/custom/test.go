package custom1

import "fmt"

var Val int	 
var val int

func PrintValue(s string)  {
	fmt.Println("Welcome",s)
	printVal("function data")
	fmt.Println("Capital Val can Explosed", Val)
	val = Val
	fmt.Println("small val can not directly exposed",val)


}

func printVal(data string) {
	fmt.Println(data)
}