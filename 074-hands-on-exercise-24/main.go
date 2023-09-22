package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("The value of x is %v and type is %T.\n", x, x)

	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("between 201 and 250")
	}

	// Checking the inclusivity and exclusivity of function rand of package math
	for n := 0; n < 10; n++ {
		fmt.Printf("Rand function output: %v.\n", rand.Intn(3))
	}
}
