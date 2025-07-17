package main

import "fmt"

func main(){
	var m map[int]string

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	m1 := make(map[int]string)
	m1[1] = "hello"
	m1[2] = "world"
	fmt.Println(m1)
	
	x, ok := m1[3]
	if !ok {
		fmt.Println("m1[3] does not exist")
	} else {
		fmt.Println("m1[3]:", x)
	}

}