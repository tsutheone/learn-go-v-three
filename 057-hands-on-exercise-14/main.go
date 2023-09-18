package main

import (
	"fmt"

	"github.com/tsutheone/puppy"
)

func main() {
	x := 42
	fmt.Println("Hello World!, the meaning of life is:", x, "hahahahahahahahahahahahahahahaha!!!!")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

	puppy.From12()
}
