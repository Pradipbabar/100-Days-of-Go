# Day 22

## Middleware and Advanced Routing

Go web frameworks such as Gin or Echo. These frameworks provide built-in support for routing and middleware integration, making it easier to handle complex web application logic and implement security features like authentication and authorization.

Here's an example of how you can use middleware for logging and authentication in a Go web application using the Gin framework:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Logging middleware
	router.Use(func(c *gin.Context) {
		fmt.Println("Request received")
		c.Next()
	})

	// Authentication middleware
	router.Use(func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")
		if authToken != "mysecrettoken" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	})

	// Secure endpoint that requires authentication
	router.GET("/secured", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You have accessed the secured endpoint"})
	})

	router.Run(":8080")
}
```

In this example, we create a Gin router and apply middleware for logging and authentication. The logging middleware logs each incoming request, and the authentication middleware checks for a specific authorization token in the request header. If the token is missing or invalid, the middleware returns an unauthorized error. We then define a secured endpoint that requires authentication to access.

By using middleware in your Go web application, you can intercept and process HTTP requests, add logging and authentication logic, and secure your routes and endpoints effectively. Middleware chaining allows you to apply multiple middleware functions in a specific order to handle complex web application logic and ensure secure request processing. You can extend this example to include more sophisticated middleware for handling various aspects of your web application, such as error handling, request validation, and data encryption.

### Sources

- <https://medium.com/@ansujain/mastering-middleware-in-go-tips-tricks-and-real-world-use-cases-79215e72b4a8>
- <https://dev.to/theghostmac/understanding-and-crafting-http-middlewares-in-go-3183>
