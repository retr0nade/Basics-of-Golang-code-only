package main

import "fmt"

func main(){
	type Power struct {
		HP int
		Torque int
	}

	type Car struct {
		Name string
		Type string
		Year int
		Power // Embedded struct
	}

	var c = Car{
		Name: "Tesla",
		Type: "Electric",
		Year: 2022,
	}

	fmt.Println("Car Name:", c.Name)
	fmt.Println("Car Type:", c.Type)
	fmt.Println("Car Year:", c.Year)
	fmt.Println("Car Details:", c)

	var c1 = Car{"Ford", "Petrol", 2021, Power{1200, 300}}

	fmt.Println("Car 1 Name:", c1.Name)
	fmt.Println("Car 1 Type:", c1.Type)
	fmt.Println("Car 1 Year:", c1.Year)
	fmt.Println("Car 1 Power HP:", c1.HP)
	fmt.Println("Car 1 Power Torque:", c1.Torque)
	fmt.Println("Car 1 Details:", c1)

	//anonymous struct
	p := struct{
		Name string
		Age int
	}{"bob", 45}

	fmt.Println("Anonymous Struct Name:", p.Name)
	fmt.Println("Anonymous Struct Age:", p.Age)
	fmt.Println("Anonymous Struct Details:", p)

}	