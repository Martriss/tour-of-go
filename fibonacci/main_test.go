package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}

	f := fibonacci()

	for i, want := range expected {
		got := f()
		if got != want {
			t.Errorf("Fibonacci sequence at position %d = %d, want %d", i, got, want)
		}
	}
}

// A benchmark test to measure performance
func BenchmarkFibonacci(b *testing.B) {
	for b.Loop() {
		f := fibonacci()
		for range 10 {
			f()
		}
	}
}
