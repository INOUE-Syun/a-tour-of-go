package main

import (
	"fmt"
	. "math"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n := 0
	return func() int {
		v := int(Round((Pow((1+Sqrt(5))/2, float64(n)) - Pow((1-Sqrt(5))/2, float64(n))) / Sqrt(5)))
		n++
		return v
	}
}

// func fibonacci() func() int {
// 	a, b, c := 0, 1, 0
// 	return func() int {
// 		a = b
// 		b = c
// 		c = a + b
// 		return c
// 	}
// }

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
