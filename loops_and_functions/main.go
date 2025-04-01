package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for range 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println("Final:   ", Sqrt(2))
	fmt.Println("Control: ", math.Sqrt(2))
}
