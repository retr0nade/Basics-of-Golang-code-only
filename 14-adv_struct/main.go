package main

import "fmt"

type Car struct {
	name string
	year int
}

func (c Car) operational(){
	if c.year > 2020 {
		fmt.Println("This car is operational.")
	} else {
		fmt.Println("This car is not operational.")
	}
}

func main(){
	c := Car{"musgoat", 2023}
	fmt.Println("Car Name:", c.name)
	fmt.Println("Car Year:", c.year)

	c.operational()

	c.year = 2019 // Update the year
	fmt.Println("Updated Car Year:", c.year)
	c.operational() // Check operational status again
}