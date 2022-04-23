package main

import (
	"fmt"
	"math/rand"
	"time"
)

var companies = []string{"AAA", "BBB", "CCC"}

func main() {
	fmt.Printf("=== Test ===\n\n")

	fmt.Printf("Spaceline        Days Trip type  Price\n")
	fmt.Println("======================================")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		num := rand.Intn(80)
		fmt.Println(num)
	}

	fmt.Println()

	fmt.Println(len(companies))

	for i := 0; i < 10; i++ {
		ind := rand.Intn(3)
		fmt.Println(companies[ind])
	}
}
