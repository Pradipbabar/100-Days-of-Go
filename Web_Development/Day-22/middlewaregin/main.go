package main

import (
 "fmt"
 "net/http"
 "time"

 "github.com/gin-gonic/gin"
)

func main() {
 r := gin.New()

 // Middleware 1 - Request validation
 r.Use(validateRequest)

 // Middleware 2 - Request tracing
 r.Use(traceRequest)

 // Middleware 3 - Security
 r.Use(secureRequest)

 // Middleware 4 - Authentication
 r.Use(authenticateRequest)

 // Login Endpoint
 r.POST("/login", login)

 if err := r.Run(":8080"); err != nil {
  fmt.Println("Failed to start server")
 }
}

func login(c *gin.Context) {
 // Authentication passed, do login logic
 c.JSON(http.StatusOK, gin.H{
  "message": "Logged in successfully",
 })
}

func validateRequest(c *gin.Context) {
 // Validate the request

 // Call the next middleware or endpoint handler
 c.Next()
}

func traceRequest(c *gin.Context) {
 // Do some request tracing before request is processed
 beforeRequest(c)

 // Call the next middleware or endpoint handler
 c.Next()

 // Do some request tracing after request is processed
 afterRequest(c)
}

func secureRequest(c *gin.Context) {
 // Add some security features to the request

 // Call the next middleware or endpoint handler
 c.Next()
}

func authenticateRequest(c *gin.Context) {
 // Authenticate the request

 // Call the next middleware or endpoint handler
 c.Next()
}

func beforeRequest(c *gin.Context) {
 start := time.Now()

 // Log the request start time
 fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)

 // Add start time to the request context
 c.Set("startTime", start)
}

func afterRequest(c *gin.Context) {
 // Get the start time from the request context
 startTime, exists := c.Get("startTime")
 if !exists {
  startTime = time.Now()
 }

 // Calculate the request duration
 duration := time.Since(startTime.(time.Time))

 // Log the request completion time and duration
 fmt.Printf("Completed %s %s in %v\n", c.Request.Method, c.Request.URL.Path, duration)
}