package testfunc

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func HardlinkSymlink() {


// 	err := os.Link("demo.txt", "demo_also.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("sym link")

// 	err = os.Symlink("demo.txt", "demo_also.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fileInfo, err := os.Lstat("demo_sym.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("file info %+v", fileInfo)

// 	err = os.Lchown("demo_sym.txt", os.Getuid(), os.Getgid())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }