package main

import (
	"fmt"
	"math/rand"
	"time"
	//	"time"
)

func main() {
	// now := time.Now()

	// fmt.Println("Time:", now)
	// fmt.Println("Seconds:", now.Unix())
	// fmt.Println("Milliseconds:", now.UnixMilli())
	// fmt.Println("Nanoseconds:", now.UnixNano())

	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Random integer
	fmt.Println(rand.Intn(100)) // 0–99

	// Random float
	fmt.Println(rand.Float64()) // 0.0–1.0
}
