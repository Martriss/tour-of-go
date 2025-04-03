package main

import (
	"fmt"
	"time"
)

func main() {
	var m SimpleMutex

	// First goroutine
	go func() {
		fmt.Println("Goroutine 1: Attempting to lock")
		m.Lock()
		fmt.Println("Goroutine 1: Locked")
		time.Sleep(2 * time.Second) // Hold the lock for 2 seconds
		fmt.Println("Goroutine 1: Unlocking")
		m.Unlock()
	}()

	time.Sleep(100 * time.Millisecond) // Give time for goroutine 1 to acquire lock

	// Second goroutine
	fmt.Println("Goroutine 2: Attempting to lock")
	m.Lock()
	fmt.Println("Goroutine 2: Locked")
	m.Unlock()
}
