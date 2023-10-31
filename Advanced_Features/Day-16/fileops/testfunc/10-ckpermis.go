package testfunc

// import (
// 	"log"
// 	"os"
// )

// func CheckFilePermission() {

// 	file, err := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
// 	if err != nil {
// 		if os.IsPermission(err) {
// 			log.Println("file write")
// 		}
// 	}
// 	file.Close()

// 	file, err = os.OpenFile("demo.txt", os.O_RDONLY, 0666)
// 	if err != nil {
// 		if os.IsPermission(err) {
// 			log.Println("file read")
// 		}
// 	}
// 	file.Close()
// }