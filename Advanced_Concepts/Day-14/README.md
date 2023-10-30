# Day 14

## File Handling in Go

In Go, file handling involves various operations such as reading from and writing to files. It also includes error handling mechanisms to manage file-related operations effectively. Here's an overview of file handling in Go:

1. **Opening a File**: To open a file, you can use the `os.Open` function. It returns a file pointer and an error if the operation fails.

   ```go
   file, err := os.Open("filename.txt")
   if err != nil {
       // handle error
   }
   defer file.Close() // Close the file at the end
   ```

2. **Reading from a File**: To read data from a file, you can use the `Read` or `ReadAll` functions from the `io/ioutil` package. Alternatively, you can use a `Scanner` for more complex reading operations.

   ```go
   data, err := ioutil.ReadAll(file)
   if err != nil {
       // handle error
   }
   fmt.Println(string(data))
   ```

3. **Writing to a File**: To write data to a file, you can use the `Write` or `WriteString` functions from the `os` package. Make sure to handle errors during the writing process.

   ```go
   data := []byte("Hello, this is some data")
   _, err := file.Write(data)
   if err != nil {
       // handle error
   }
   ```

4. **Error Handling**: Go provides various error handling mechanisms, including checking for `nil` values and using the `error` interface. You can use conditional statements or the `if err != nil` pattern to handle file-related errors.

   ```go
   if err != nil {
       fmt.Println("Error:", err)
   }
   ```

Proper error handling is crucial when working with file operations in Go. It ensures that your program can gracefully handle any errors that may occur during file reading, writing, or other file-related operations.
