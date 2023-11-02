# Day 21

## Working with Templates in Go

To use templates for generating dynamic HTML content in Go, you can leverage the `html/template` package, which provides a powerful and secure way to create HTML templates. Here's a basic example of how to use the `html/template` package:

```go
package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Define a simple template
	const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Hello!</title>
</head>
<body>
    <h1>Hello, {{.Name}}!</h1>
    <p>You are {{.Age}} years old.</p>
</body>
</html>
`

	// Create a new template and parse the template string
	t := template.Must(template.New("person").Parse(tmpl))

	// Define a person
	person := Person{Name: "John Doe", Age: 30}

	// Execute the template with the person data and write the output to the standard output
	err := t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
```

In this example, we create a simple HTML template that includes dynamic data fields `{{.Name}}` and `{{.Age}}`. We define a `Person` struct with the corresponding fields, and then use the `template.New` and `template.Parse` functions to create and parse the template. Finally, we execute the template with the `Person` data and write the output to the standard output.

The `html/template` package provides a powerful yet safe way to generate HTML content, ensuring that the output is properly escaped to prevent cross-site scripting (XSS) attacks. You can use this package to create more complex templates, iterate over data structures, and include conditional logic within your HTML templates.
