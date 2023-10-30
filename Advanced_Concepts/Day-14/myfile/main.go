package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - raw data "

	file, err := os.Create("./Advanced_Concepts/Day-14/myfile/testfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./Advanced_Concepts/Day-14/myfile/testfile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}