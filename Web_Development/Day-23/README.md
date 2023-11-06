# Day 23

## Working with Databases in Go

To work with databases in Go, you can utilize the `database/sql` package, which provides a generic interface for working with SQL databases, as well as popular ORM (Object Relational Mapping) libraries like GORM or XORM. These tools enable you to establish connections to various databases, execute SQL queries, and perform CRUD operations seamlessly within your Go applications. Here's a simplified example demonstrating how to connect to a PostgreSQL database and perform basic CRUD operations using the GORM library:

```go
package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to the database")

	// Auto Migrate the User struct
	db.AutoMigrate(&User{})

	// Create
	user := User{Name: "John Doe", Age: 30}
	result := db.Create(&user)
	if result.Error != nil {
		panic("Failed to create user")
	}
	fmt.Println("Created user:", user)

	// Read
	var retrievedUser User
	db.First(&retrievedUser, user.ID)
	fmt.Println("Retrieved user:", retrievedUser)

	// Update
	db.Model(&retrievedUser).Update("Age", 40)
	fmt.Println("Updated user:", retrievedUser)

	// Delete
	db.Delete(&retrievedUser)
	fmt.Println("Deleted user:", retrievedUser)

	// Close the database connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get SQL DB from GORM")
	}
	sqlDB.Close()
}
```

This example demonstrates how to connect to a PostgreSQL database using GORM, create a `User` model, and perform basic CRUD operations such as creating, reading, updating, and deleting records. You can adjust the database connection settings and model structure based on your specific use case and database configuration. Understanding these techniques can empower you to effectively manage data in your Go applications and interact with different types of databases efficiently.

### Sources

- <https://earthly.dev/blog/using-gorm-go/>
- <https://dev.to/techschoolguru/generate-crud-golang-code-from-sql-and-compare-db-sql-gorm-sqlx-sqlc-560j>
- <https://medium.com/mercadolibre-tech/go-language-relational-databases-and-orms-682a5fd3bbb6>
