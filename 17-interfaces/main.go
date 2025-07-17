package main

import "fmt"

type Car interface {
	startengine() string
	drive() string
}

type Suv struct {
	name string
	year int
}

type Sedan struct {
	name string
	month int
}

func (s Suv) startengine() string {
	return fmt.Sprintf("Starting engine of SUV %s from %d", s.name, s.year)
}

func (s Sedan) startengine() string {
	return fmt.Sprintf("Starting engine of Sedan %s from month %d", s.name, s.month)
}

func (s Suv) drive() string {
	return fmt.Sprintf("Driving SUV %s from %d", s.name, s.year)
}	

func (s Sedan) drive() string {
	return fmt.Sprintf("Driving Sedan %s from month %d", s.name, s.month)
}

func definecar(c Car){
	fmt.Println(c.startengine())
	fmt.Println(c.drive())
}

func main() {
	var car Car
	car = Suv{name: "Toyota", year: 2020}
	definecar(car)
	car = Sedan{name: "Honda", month: 5}
	definecar(car)
}

