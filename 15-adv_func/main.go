package main

import (
	"fmt"
)

func operation(x,y int, op func(x,y int)int) int{
	return op(x,y)
}

func add(x,y int) int{
	return x + y
}

func multiply(x,y int) int{
	return x * y
}


func multiplier(factor int) func(int)int{
	return func(x int) int {
		return x * factor
	}
}


func main(){
	//1. functions as input

	fmt.Println("Addition:", operation(5, 3, add))       // Using add function
	fmt.Println("Multiplication:", operation(5, 3, multiply)) // Using multiply function


	//2. functions as output
	double := multiplier(2) // Create a function that doubles its input
	fmt.Println("Double of 5:", double(5)) // Call the returned function with 5
	triple := multiplier(3) // Create a function that triples its input
	fmt.Println("Triple of 5:", triple(5)) // Call the returned function with 5

}