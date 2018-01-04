package main

import (
	"fmt"
	"math"
)

const small = 0.000001

func Sqrt(x float64) {
	starts := [4]float64{1.0, x, x / 2, x / 4}
	for i := 0; i < len(starts); i++ {
		start := starts[i]
		result, tries := sqrt(x, start)
		fmt.Printf("num=%.2f, sqrt=%.2f, start=%.2f, tries=%d\n", x, result, start, tries)
	}
}

func sqrt(x, z float64) (float64, int) {
	p := math.MaxFloat64
	tries := 0
	for math.Abs(z-p) > small {
		p = z
		z -= (z*z - x) / (2 * z)
		tries++
	}
	return z, tries
}

func main() {
	Sqrt(4)
	Sqrt(25)
	Sqrt(33)
}

