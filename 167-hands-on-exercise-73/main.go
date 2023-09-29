// Package main: Wrapper function.
package main

import (
	"fmt"
	"time"
)

func main() {
	logWorkTime(doWork)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

func logWorkTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

/*
 */
