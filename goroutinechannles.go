package main

import (
	"fmt"
	"sync"
)

func main() {
	var wt sync.WaitGroup
	wt.Add(1)

	// Create a channel to receive the boolean
	ch := make(chan bool)

	// Start Test as a goroutine directly
	go TestWithChannel(&wt, ch)

	// Receive the value from the channel
	result := <-ch
	fmt.Println("Returned value:", result)

	wt.Wait()
}

// New version of Test that takes a channel to send result
func TestWithChannel(wt *sync.WaitGroup, ch chan bool) {
	defer wt.Done()
	ch <- true // send the return value into the channel
}