package main

import (
	"fmt"
)

type MotorVehicle interface {
	Milage() float64
	
}

type BMW struct{
	distance float64
	fuel float64
	avragespeed string
}


func (b BMW) Milage() float64{
	return b.distance/b.fuel
}


func totalMileage(m MotorVehicle){
	au := m.(BMW)
	fmt.Printf(au.avragespeed)
}

func main() {
	fmt.Println("hello")
	b1 := BMW{
		distance: 167.9,
		fuel: 36,
		avragespeed: "58",
	}

	totalMileage(b1)
}