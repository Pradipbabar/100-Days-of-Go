package testfunc

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func SeekPositionInFile() {

// 	/*
// 	   Seek Positions in File
// 	*/

// 	file, _ := os.Open("demo.txt")
// 	defer file.Close()

// 	// How many bytes should the offset move?
// 	// Offset can be negative(-) or positive(+)
// 	var offset int64 = 5

// 	// Where's the reference point for offset?
// 	// 0 = Beginning of file
// 	// 1 = Current position
// 	// 2 = End of file
// 	var whence int = 0
// 	newPosition, err := file.Seek(offset, whence)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Just moved to 5: ", newPosition)

// 	// Go back 2 bytes from current position
// 	newPosition, err = file.Seek(-2, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Just moved back two: ", newPosition)

// 	// Find the current position by hetting the return value from Seek after moving 0 bytes.
// 	currentPosition, err := file.Seek(0, 1)
// 	fmt.Println("Current position: ", currentPosition)

// 	// Go to the beginning of the file.
// 	newPosition, err = file.Seek(0, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Position after seeking 0,0: ", newPosition)
// }