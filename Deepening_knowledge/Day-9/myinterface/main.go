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

type Audi struct{
	distance float64
	fuel float64
	
}

func (b BMW) Milage() float64{
	return b.distance/b.fuel
}

func (a Audi) Milage() float64{
	return a.distance/a.fuel
}

func totalMileage(m []MotorVehicle){
	tm := 0.0
	for _, v:= range m {
		tm = tm + v.Milage()
	}
	fmt.Printf("Total Mileage per Month is : %f km/l",tm)
}

func main() {
	fmt.Println("hello")
	b1 := BMW{
		distance: 167.9,
		fuel: 36,
		avragespeed: "58",
	}

	a1 := Audi{
		distance:152.9,
		fuel: 30,

	}

	person := []MotorVehicle{b1,a1}
	totalMileage(person)
}