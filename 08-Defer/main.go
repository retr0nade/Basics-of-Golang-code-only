package main

import "fmt"

func hw(name string){
	fmt.Println("Hello", name)
}

func main(){
	defer hw("World") //defer statement will execute after the main function completes
	defer hw("Golang") //this will execute before the above defer statement
	hw("Shreyas") //this will execute first
}
