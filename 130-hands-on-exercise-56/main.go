package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "pepe",
		friends: map[string]int{
			"manolo": 56,
			"pedro":  62,
			"carlos": 39,
		},
		favDrinks: []string{"beer", "wine", "sangria"},
	}

	fmt.Println(p1.first)
	for k, v := range p1.friends {
		fmt.Println("friend:", k, v)
	}
	for i, v := range p1.favDrinks {
		fmt.Println("drink:", i, v)
	}
}

/*
Create and use an anonymous struct with these fields:
first     string
friends   map[string]int
favDrinks []string
Print things.
*/
