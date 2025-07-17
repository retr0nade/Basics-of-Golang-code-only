package main

import(
	"fmt"
	"sync"
	"time"
)

func worker(i int, delay int, wg *sync.WaitGroup){
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	fmt.Println("Worker no:",i," Processing started")
	time.Sleep(time.Duration(delay) *time.Second) // Simulate some work
	fmt.Println("Worker Processing Completed")
}

func main(){
	var wg sync.WaitGroup

	//for i:=0; i<3; i++{
	//	wg.Add(1) // Increment the WaitGroup counter
	//	go worker(i+1, &wg) // Start a new goroutine
	//}

	wg.Add(5) // Increment the WaitGroup counter for 5 workers
	go worker(1,2, &wg) // Start first worker
	go worker(2,3, &wg) // Start second worker
	go worker(3,1, &wg) // Start third worker
	go worker(4,4, &wg) // Start fourth worker
	go worker(5,2, &wg) // Start fifth worker


	wg.Wait() // Wait for all goroutines to finish
}