package main

import "fmt"

const z = 0

func main(){
	const a = 10
	const pi float64 = 3.14
	fmt.Println(a, pi, z)

	const(
		L = 5
		B = 10
		A = L*B
	)
	fmt.Println("Area of rectangle:", A)

	const (
		zero = iota //0
		one //1
		two //2
	)
	fmt.Println("zero:", zero, "one:", one, "two:", two)
}


