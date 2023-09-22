package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println("This is x:", x)
	fmt.Println("This is y:", y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4.")
	case x > 6 && y > 6:
		fmt.Println("x and y are both greater than 6.")
	case x >= 4 && x <= 6:
		fmt.Println("x is between 4 and 6 both inclusive.")
	case y != 5:
		fmt.Println("y is not 5.")
	default:
		fmt.Println("None of the conditions were met, defaults.")
	}
}
