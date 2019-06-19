package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Set the seed to get random number
	// If you don't, rand returns same value
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
