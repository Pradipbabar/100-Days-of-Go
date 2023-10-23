// Golang program to illustrate how to pass
// a function as an argument to another
// function
package main

import (
	"fmt"
	"math"
)

// Volume function takes
// a function as a argument
func Sphere(vol func(radius float64) float64) {

	fmt.Println("Volume of Sphere is:", vol(3))
}

// Main Function
func main() {

	volume_of_sphere := func(radius float64) float64 {
		result := 4 / 3 * math.Pi * radius * radius * radius
		return result
	}

	// Passing function as an
	// argument in Volume function
	Sphere(volume_of_sphere)
}
