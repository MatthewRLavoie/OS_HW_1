package main

import (
	"fmt" // Print results
	"time" // Time delay to prevent incorrect order of goroutine execution
)

// Producer goroutine
func producer(ch chan int) {
	//increamental counter
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: %d\n", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch) // Close channel to signal completion
}

//consumer goroutine
func consumer(ch chan int) {
	for num := range ch {
		fmt.Printf("Consumer: %d\n", num)
	}
}

//main execution
func main() {
	ch := make(chan int) // Unbuffered channel for communication
	go producer(ch)
	consumer(ch) // Running consumer in the main goroutine
}
