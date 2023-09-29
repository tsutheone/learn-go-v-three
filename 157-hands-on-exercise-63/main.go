package main

import "fmt"

func main() {
	result := factorial(5)
	fmt.Println(result)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

/*
● Create an example of a test in go
*/
