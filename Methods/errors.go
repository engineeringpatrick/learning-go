package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// now ErrNegativeSqrt implements "error" interface
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := 1.0
	for oldZ := x; math.Abs(z-oldZ) > 0.000000001; {
		oldZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
