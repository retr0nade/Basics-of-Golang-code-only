package main

import (
	"fmt"
	"time"
)

func send(ch chan <- int){
	fmt.Println("Sending value to channel")
	time.Sleep(2 * time.Second)
	ch <- 20 // Send value to the channel
}

func receive(ch <- chan int){
	fmt.Println("Receiving value from channel")
	value := <-ch // Receive value from the channel
	fmt.Println("Received value:", value)
}

func main(){
	ch := make(chan int) // Create a channel of type int
	
	go func(){
		time.Sleep(1 * time.Second) // Simulate some work before sending
		ch <- 20
	}()

	receive(ch)
}