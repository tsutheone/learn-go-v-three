package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("key %v \t value %v.\n", k, v)
	}

	age := m["James"]
	fmt.Println(age)

	qcheck := m["Q"]
	fmt.Printf("u %v \t type %T.\n", qcheck, qcheck)

	if _, ok := m["Q"]; !ok {
		fmt.Println("No value stored in key Q in map m")
		u := !ok
		fmt.Printf("u %v \t type %T.\n", u, u)
	}
}
