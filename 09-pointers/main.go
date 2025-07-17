package main

import "fmt"

func incr(n *int) int{
	return *n + 1	
}

func main(){
	var p *int
	var x int = 10
	p = &x
	fmt.Println("Value of x: ", p)
	fmt.Println("Address of x: ", *p)

	*p = incr(p)
	fmt.Println("Value of x after increment: ", *p)

	pp := &p
	fmt.Println("Value of x: ", **pp)
}