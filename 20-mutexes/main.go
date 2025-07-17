package main

import(
	"fmt"
	"sync"
)

var(
	counter int
	mu sync.Mutex // Mutex to protect the counter
)

func increment(wg *sync.WaitGroup){
	defer wg.Done() // Notify the WaitGroup that this goroutine is done
	mu.Lock()
	counter++
	mu.Unlock()
}

func main(){
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		wg.Add(1)
		go increment(&wg) // Start a new goroutine to increment the counter
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter value:", counter) // Print the final value of the counter
}