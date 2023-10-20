# Day 3

## Control Structures in Go

Understanding control structures is crucial for writing effective and dynamic programs. In Go, the `if`, `for`, and `switch` statements are fundamental control structures. Let's explore each of these control structures and then look at some simple examples:

### 1. if Statement:

The `if` statement allows you to execute a block of code if a specified condition is true. It has a simple syntax:

```go
if condition {
    // Code to be executed if condition is true
} else if anotherCondition {
    // Code to be executed if anotherCondition is true
} else {
    // Code to be executed if none of the conditions are true
}
```

### 2. for Loop:

The `for` loop in Go is used to repeatedly execute a block of code as long as a specified condition is true. There are different ways to use the `for` loop:

- **Basic for loop**:

```go
for i := 0; i < 5; i++ {
    // Code to be executed
}
```

- **For loop as a while loop**:

```go
for condition {
    // Code to be executed
}
```

- **Infinite loop**:

```go
for {
    // Code to be executed indefinitely
}
```

### 3. switch Statement:

The `switch` statement is used to perform different actions based on different conditions. It's a cleaner way to write multiple `if-else` statements. It has the following syntax:

```go
switch variable {
case condition1:
    // Code to be executed if variable matches condition1
case condition2:
    // Code to be executed if variable matches condition2
default:
    // Code to be executed if none of the conditions match
}
```

### Examples:

Let's use these control structures in some simple examples:

1. Example of an `if` statement:

```go
package main

import "fmt"

func main() {
    num := 10
    if num > 5 {
        fmt.Println("Number is greater than 5")
    } else {
        fmt.Println("Number is less than or equal to 5")
    }
}
```

2. Example of a `for` loop:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

3. Example of a `switch` statement:

```go
package main

import "fmt"

func main() {
    num := 2
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    default:
        fmt.Println("Other number")
    }
}
```

