package testfunc

import (
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadData() {

	/*
	   Downloading a File Over HTTP
	*/

	// Create output file
	newFile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request 
	url := "https://raw.githubusercontent.com/Pradipbabar/100-Days-of-Go/main/README.md"
	response, err := http.Get(url)
	defer response.Body.Close()

	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}