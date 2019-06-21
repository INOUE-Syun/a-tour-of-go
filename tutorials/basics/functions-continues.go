package main

import "fmt"

// If parameters have same type, It can omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
