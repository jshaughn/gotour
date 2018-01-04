package main

import (
	"fmt"
	"math"
)

const small = 0.000001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) String() string {
	return fmt.Sprint(float64(e))
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %2.0f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	starts := [4]float64{1.0, x, x / 2, x / 4}
	var result float64
	var tries int
	for i := 0; i < len(starts); i++ {
		start := starts[i]
		result, tries = sqrt(x, start)
		fmt.Printf("num=%.2f, sqrt=%.2f, start=%.2f, tries=%d\n", x, result, start, tries)
	}

	return result, nil
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
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

