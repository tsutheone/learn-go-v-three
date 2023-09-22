package main

import "fmt"

func main() {
	for x := 0; x <= 42; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
