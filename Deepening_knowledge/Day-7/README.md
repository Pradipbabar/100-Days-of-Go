# Day 7

## Pointers, Structs and Methods 

### Pointers

In Go, a pointer is a variable that stores the memory address of another variable. Using pointers allows you to manage and control the memory directly, providing a way to pass references to values and data structures rather than passing a copy of the data. Understanding pointers in Go is essential for managing memory, creating data structures, and optimizing performance.

### Declaring Pointers in Go:

You can declare a pointer using the `*` operator. For example:

```go
var ptr *int
```

This declaration indicates that `ptr` is a pointer to a variable of type `int`.

### Using Pointers in Go:

You can use pointers to access and modify the value of the variable they point to. The `&` operator is used to get the memory address of a variable, and the `*` operator is used to dereference a pointer, which means accessing the value stored at the memory address the pointer is pointing to. For example:

```go
var num int = 42
var ptr *int = &num
fmt.Println(*ptr) // Prints the value stored at the memory address pointed to by ptr
```

### Advanced Concepts:

1. **Pointer Arithmetic**:
   - Unlike some programming languages, Go does not support pointer arithmetic. This is done intentionally for safety and to avoid certain types of bugs common in other languages.

2. **Passing Pointers to Functions**:
   - Go allows you to pass pointers to functions, enabling functions to modify the original values directly. This can be particularly useful when you want to avoid copying large data structures.

3. **Pointers to Structs**:
   - Pointers to structs are frequently used in Go to avoid copying structs when passing them to functions or methods. This allows for more efficient memory usage and prevents unnecessary data duplication.

4. **Pointer to Pointer**:
   - Go also supports pointers to pointers, which are used in scenarios where you need to modify a pointer itself.

### Structs

In Go, a struct is a composite data type that allows you to compose together values of different types into a single entity. It is used to create user-defined types which can contain named fields. Structs are particularly useful for defining complex data types and are widely used in Go for various purposes, including representing objects, data records, and more.

### Declaring and Defining a Struct in Go:

Here's how you can declare and define a struct in Go:

```go
// Defining a struct type
type Person struct {
    name string
    age  int
}
```

In this example, `Person` is the name of the struct, and it has two fields: `name` of type `string` and `age` of type `int`.

### Creating Instances of a Struct:

You can create instances of a struct using the following syntax:

```go
// Creating an instance of the struct
var p Person
p.name = "John"
p.age = 30
```

### Struct Initialization:

You can also initialize a struct using a struct literal during declaration:

```go
p := Person{name: "John", age: 30}
```

### Accessing Struct Fields:

You can access the fields of a struct using the dot (`.`) operator:

```go
fmt.Println(p.name) // Output: John
fmt.Println(p.age)  // Output: 30
```

### Structs with Methods:

Go allows you to associate methods with a struct type. This is similar to adding methods to a class in object-oriented programming. Here's an example:

```go
func (p Person) introduce() {
    fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.name, p.age)
}
```

### Anonymous Structs:

Go also allows you to create anonymous structs, which are useful when you need to create a struct for a one-time use:

```go
point := struct {
    x, y int
}{10, 15}
```

### Embedded Structs:

Go supports embedded structs, which allows a struct type to include the fields of another struct type without explicitly declaring them. This is similar to inheritance in object-oriented programming.

```go
type Contact struct {
    email, phone string
}

type Employee struct {
    Person
    Contact
    employeeID int
}
```

### Methods 

In Go, methods are functions that operate on a specific type. They are associated with the type they operate on and enable you to define behavior for that type. Methods in Go are similar to methods in other programming languages but have some distinctive features that make them particularly powerful and flexible.

### Syntax for Declaring a Method in Go:

You can declare a method for a type using the following syntax:

```go
func (t Type) methodName(parameterList) {
    // Method implementation
}
```

where:
- `t` is the receiver, which can be a value or a pointer.
- `Type` is the type the method operates on.
- `methodName` is the name of the method.

### Methods with Value Receivers and Pointer Receivers:

In Go, methods can have either value receivers or pointer receivers. Value receivers operate on a copy of the data, while pointer receivers operate on the actual data. Here's an example of both types of methods:

```go
// Value receiver method
func (p Person) introduce() {
    fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.name, p.age)
}

// Pointer receiver method
func (p *Person) changeName(newName string) {
    p.name = newName
}
```

### Method Invocation in Go:

You can call methods on an instance of a type using the dot (`.`) operator. Here's an example:

```go
p := Person{name: "John", age: 30}
p.introduce() // Calling the value receiver method
p.changeName("Mike") // Calling the pointer receiver method
```

### Method Sets in Go:

A method set is the set of methods that can be called on a specific type. It includes methods with both value receivers and pointer receivers. Understanding method sets is essential when working with interfaces in Go.

### Important Notes:

- Go does not support method overloading or method overriding, but it does support method receivers of the same name on different types.
- You can define methods for any named type in Go, not just structs.

Methods in Go enable you to add behavior to your types and provide a convenient way to organize and encapsulate functionality. They play a crucial role in defining the behavior of user-defined types and are widely used in Go programming for creating efficient and expressive code.
