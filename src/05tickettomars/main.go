package main

import (
	"fmt"
	"math/rand"
	"time"
)

var companies = []string{"SpaceX", "Space Adventures", "Virgin Galactic"}
var tripType = []string{"One-way", "Round-trip"}

const distance = 62100000

func main() {
	fmt.Printf("=== Test ===\n\n")

	fmt.Printf("Spaceline        Days Trip type  Price\n")
	fmt.Println("======================================")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		num := rand.Intn(80)
		fmt.Println("test:", num)
	}

	fmt.Println()

	// fmt.Println(len(companies))

	for i := 0; i < 10; i++ {
		ind := rand.Intn(3)
		fmt.Println("comp:", companies[ind])
	}

	fmt.Println()

	for i := 0; i < 10; i++ {
		speed := rand.Intn(15) + 16
		fmt.Println("speed:", speed)
	}
}
