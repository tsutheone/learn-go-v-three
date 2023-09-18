package main

import (
	"fmt"
)

const combinations = 139838160
const eurNumbers = 5
const eurStars = 2

var talk = "Euromillion consist of:"

func main() {

	eurSize := eurNumbers + eurStars
	chances := (1.0 / combinations) * 100

	fmt.Println(talk, eurNumbers, "numbers and", eurStars, "stars.")
	fmt.Println("For a total size of", eurSize, "numbers.")
	fmt.Println("I have a chance of winning of", chances, "%")
	fmt.Println("")
	fmt.Printf("The value of combinations is %v and the type is %T\n", combinations, combinations)
	fmt.Printf("The value of eurNumbers is %v and the type is %T\n", eurNumbers, eurNumbers)
	fmt.Printf("The value of eurStars is %v and the type is %T\n", eurStars, eurStars)
	fmt.Printf("The value of talk is \"%v\" and the type is %T\n", talk, talk)
	fmt.Printf("The value of eurSize is %v and the type is %T\n", eurSize, eurSize)
	fmt.Printf("The value of chances is %v and the type is %T\n", chances, chances)

}
