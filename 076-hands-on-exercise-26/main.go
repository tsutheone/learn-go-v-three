package main

import (
	"fmt"
	"math/rand"
)

// The init function is a niladic function, that means that never takes and argument. Admits no arguments.
func init() {
	fmt.Println("This is where the initialization for my program occurs.")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v and type is %T.\n", x, x)

	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("shouldn't be here")
	}

	// Checking the inclusivity and exclusivity of function rand of package math
	for n := 0; n < 10; n++ {
		fmt.Printf("Rand function output: %v.\n", rand.Intn(3))
	}
}
