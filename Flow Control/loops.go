package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for oldZ := x; math.Abs(z-oldZ) > 0.000000001; {
		oldZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main_loops() {
	fmt.Println(Sqrt(25))
	fmt.Println()
	fmt.Println(math.Sqrt(25))
}
