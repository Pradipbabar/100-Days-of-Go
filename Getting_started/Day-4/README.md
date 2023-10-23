# Day 4

## Functions

In Go, functions are a fundamental building block of any program. They allow you to break down your code into smaller, reusable pieces, making your code more organized and easier to maintain. Here's how you can define and call functions in Go:

### Defining Functions in Go:

The general syntax for defining a function in Go is as follows:

```go
func functionName(parameterName parameterType) returnType {
    // Function body
    // Statements to be executed
    return someValue
}
```

In the function definition:
- `func` is the keyword used to define a function.
- `functionName` is the name you give to your function.
- `parameterName` and `parameterType` are the input parameters that the function accepts.
- `returnType` is the type of value the function returns.

### Calling Functions in Go:

To call a function in Go, use the function name followed by parentheses. If the function has parameters, you need to provide the values for those parameters inside the parentheses.

### Example of Function Definition and Function Call:

```go
package main

import "fmt"

// Function definition
func add(x int, y int) int {
    return x + y
}

func main() {
    // Function call
    result := add(5, 3)
    fmt.Println("Result:", result)
}
```

### Function Parameters and Return Types:

1. **Function Parameters**:
   - Go supports passing parameters by value. You can pass various types of parameters, such as integers, floats, strings, arrays, slices, maps, structs, pointers, and even functions.
   - Parameters can be of the same type or different types. For the same type, you can specify the type once at the end.

2. **Return Types**:
   - Functions can return one or more values. You can specify the return type(s) after the list of parameters.
   - If a function doesn't explicitly return a value, it is considered to return `nil` or the zero value of the specified type.

### Example of a Function with Multiple Return Values:

```go
package main

import "fmt"

// Function with multiple return values
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("Cannot divide by zero")
    }
    return x / y, nil
}

func main() {
    result, err := divide(6, 3)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

Understanding how to define and call functions, and how to handle parameters and return types, will allow you to write modular and maintainable code in Go.