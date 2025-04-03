package main

import (
	"runtime"
	"sync/atomic"
)

// SimpleMutex is a basic implementation of a mutex
type SimpleMutex struct {
	locked int32 // Using int32 for atomic operations
}

// Lock acquires the mutex
func (m *SimpleMutex) Lock() {
	// Try to atomically change locked from 0 to 1
	for !atomic.CompareAndSwapInt32(&m.locked, 0, 1) {
		// If we couldn't acquire the lock, yield to other goroutines
		runtime.Gosched()
		// You could also add a small sleep here to reduce CPU usage
		// time.Sleep(1 * time.Microsecond)
	}
}

// Unlock releases the mutex
func (m *SimpleMutex) Unlock() {
	// Set locked back to 0
	atomic.StoreInt32(&m.locked, 0)
}
