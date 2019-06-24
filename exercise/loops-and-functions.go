package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2

	// 10 iterates.
	// for i := 0; i < 10; i++ {
	// 	z -= (z*z - x) / (2 * z)
	// }

	// end loop when value not change (or changed very small)
	var prev float64
	for {
		if math.Abs(z-prev) < 1e-10 {
			break
		}
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	x := float64(5)
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
