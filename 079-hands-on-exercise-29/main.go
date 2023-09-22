package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rarec := 0

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("iteration %v \t x is %v, y is %v \t", i, x, y)
		switch {
		case x < 4 && y < 4:
			fmt.Println("both less than 4.")
		case x > 6 && y > 6:
			fmt.Println("oth greater than 6.")
		case x >= 4 && x <= 6:
			fmt.Println("x is between 4 and 6 both inclusive.")
		case y != 5:
			fmt.Println("y is not 5.")
		default:
			fmt.Println("None of the conditions were met, defaults.")
			rarec++
		}
	}
	fmt.Printf("default appeared: %v times out of 100\n", rarec)
}
