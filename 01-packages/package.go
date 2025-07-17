package main

import ("fmt"
		"packages/math"
)

func main(){
	fmt.Println("My First Package")
	fmt.Println("Addition using math package: ", math.Add(5,3))
	fmt.Println("Value of X from math package:", math.X)
}