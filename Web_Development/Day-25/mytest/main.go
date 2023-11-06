package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func HomepageHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message":"Welcome to the Tech Company listing API with Golang"})
}

func main() {
    router := gin.Default()
    router.GET("/", HomepageHandler)
    router.GET("/companies", GetCompaniesHandler)
    router.POST("/company", NewCompanyHandler)
    router.PUT("/company/:id", UpdateCompanyHandler)
    router.DELETE("/company/:id", DeleteCompanyHandler)
    router.Run()
}