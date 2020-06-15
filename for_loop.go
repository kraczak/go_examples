package main

import (
	"fmt"
	"math"
)

func mySqrt(x float64) float64 {
	var z float64 = 1
	var n int = 10
	for i := 0; i < n; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	for k := 0; k <= 625; k += 121 {
		real := math.Sqrt(float64(k))
		mine := mySqrt(float64(k))
		fmt.Printf("%v --> Real sqrt: %v, Mine sqrt: %v\n", k, real, mine)
	}
}
