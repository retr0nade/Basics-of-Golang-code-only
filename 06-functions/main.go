package main

import "fmt"

func HW(){
	fmt.Println("Hello World!")
}

func sum(a, b int) int{
	return a+b
}

//variadic function 
// allows you to pass a variable number of arguments
// the ... before the type indicates that it can take multiple values of that type
func add(num ...int)int{
	sum := 0
	for i:= range num{
		sum += num[i]
	}
	return sum
}



func main(){
	HW()
	c := sum(2,3)
	fmt.Println("Sum is: ", c)

	one := add(1,2,3,4,5)
	two := add(6,7,8,9,10, 11, 12, 13, 14, 15)
	fmt.Println("Sum of one is: ", one)
	fmt.Println("Sum of two is: ", two)

	name := func(n string) string{
		return "Hello " + n
	}

	fmt.Println(name("WORLD"))
}