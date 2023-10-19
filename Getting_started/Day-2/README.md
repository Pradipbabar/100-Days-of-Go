# Variable And Data Types

## Variables
In Go, variable data types are used to specify the kind of data that a particular variable can store. The basic data types in Go include:

1. **Numeric Types:** These include integers (`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`) and floating-point numbers (`float32`, `float64`, `complex64`, `complex128`).

2. **Boolean Type:** Represents true or false values, declared as `bool`.

3. **String Type:** Used to store text data, declared as `string`.

4. **Derived Types:** These include pointers, arrays, structures, and more complex types like slices, maps, and channels.

As for inputs in Go, you can use the `fmt` package to handle input and output. The `fmt.Scan` function is commonly used to read input from the standard input (keyboard) into variables. Additionally, you can use the `os` package to work with command-line arguments passed to the program.

Here's an example of using `fmt.Scan` to get user input:

```go
package main

import "fmt"

func main() {
    var input string
    fmt.Print("Enter text: ")
    fmt.Scanln(&input)
    fmt.Println("You entered:", input)
}
```


### Source

- <https://go.dev/tour/basics/11>
- <https://www.w3schools.com/go/go_data_types.php>
