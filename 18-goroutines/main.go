package main

import (
	"fmt"
	"time"
	"runtime"
)

func printNumbers(){
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Another goroutine")
}

func main(){
	runtime.GOMAXPROCS(8)

	go printNumbers() // Start a new goroutine
	
	time.Sleep(10 * time.Second)
	fmt.Println("Main goroutine is running")
	// doesnt wait for the goroutine to finish
}