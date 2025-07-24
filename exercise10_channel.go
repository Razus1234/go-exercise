package main

import (
	"fmt"
	"sync"
)

func exercise10_channel() {
	fmt.Println("✅ Exercise 10: สร้าง Channel")
	var wg sync.WaitGroup
	var receiverWg sync.WaitGroup
	ch := make(chan string)

	// Receiver goroutine
	receiverWg.Add(1)
	go func() {
		defer receiverWg.Done()
		for msg := range ch {
			fmt.Println("Received message from channel:", msg)
		}
	}()

	// Sender goroutine 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Hello from channel!"
		fmt.Println("Message sent to channel.")
	}()

	// Sender goroutine 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Hi David!"
		fmt.Println("Message sent to channel.")
	}()

	// Wait for senders to finish, then close the channel
	go func() {
		wg.Wait() // wait for senders only
		close(ch) // close channel after all sends are done
	}()

	receiverWg.Wait() // wait for receiver to finish
	fmt.Println("Channel operation completed.")
}
