package testfunc

// import (
// 	"log"
// 	"os"
// 	"time"
// )

// func ChangeFilePermission() {


// 	err := os.Chmod("demo.txt", 0777)
// 	if err != nil {
// 		log.Println(err)

// 	}

// 	err = os.Chown("demo.txt", os.Getuid(), os.Getgid())
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	twoDaysFromNow := time.Now().Add(48 * time.Hour)
// 	lastAccessTime := twoDaysFromNow
// 	lastModifyTime := twoDaysFromNow
// 	err = os.Chtimes("demo.txt", lastAccessTime, lastModifyTime)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }