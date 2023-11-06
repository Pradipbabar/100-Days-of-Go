package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"uniqueIndex"`
	LastName  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"not null"`
	Country   string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

// func UpdateUser() {
// 	// Retrieve the record you want to update
// 	var user User
// 	result := db.First(&user, 1)
// 	if result.Error != nil {
// 		panic("failed to retrieve user: " + result.Error.Error())
// 	}

// 	// Modify the attributes of the retrieved record, in this case, \
// 	// the first three columns
// 	user.FirstName = "Agnes"
// 	user.LastName = "Doe"
// 	user.Email = "agnesdoe@example.com"

// 	// Save the changes back to the database
// 	result = db.Save(&user)
// 	if result.Error != nil {
// 		panic("failed to update user: " + result.Error.Error())
// 	}

// 	fmt.Println("User updated successfully")
// }

// func DeletUser() {
// 	var user User
// 	result := db.First(&user)
// 	if result.Error != nil {
// 		panic("failed to retrieve user: " + result.Error.Error())
// 	}

// 	result = db.Delete(&user)
// 	if result.Error != nil {
// 		panic("failed to delete user: " + result.Error.Error())
// 	} else if result.RowsAffected == 0 {
// 		panic("no user record was deleted")
// 	} else {
// 		fmt.Println("User record deleted successfully")
// 	}
// }

func main() {
	//Create a new Postgresql database connection
	dsn := "host=<your_host> user=<your_user> password=<your_password> dbname=<your_dbname> port=<your_port>"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}
	newUser := User{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "janedoe@gmail.com",
		Country:   "Spain",
		Role:      "Chef",
		Age:       30,
	}

	// ... Create a new user record...
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("failed to create user: " + result.Error.Error())
	}
	// ... Handle successful creation ...
	fmt.Printf("New user %s %s was created successfully!\\n", newUser.FirstName, newUser.LastName)
	// Retrive User
	var users []User
	result1 := db.Where("FirstName = ?", "Jane").Where("Country = ?", "Spain").Find(&users)
	if result1.Error != nil {
		panic("failed to retrieve users: " + result1.Error.Error())
	}

	// Use the user records
	for _, user := range users {
		fmt.Printf(" Name: %s %s, Email: %s\n", user.FirstName, user.LastName, user.Email)
	}
}
