package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producer: %d\n", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch) // Close channel to signal completion
}

func consumer(ch chan int) {
	for num := range ch {
		fmt.Printf("Consumer: %d\n", num)
	}
}

func main() {
	ch := make(chan int) // Unbuffered channel for communication
	go producer(ch)
	consumer(ch) // Running consumer in the main Goroutine
}
