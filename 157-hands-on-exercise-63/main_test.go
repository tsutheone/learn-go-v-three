package main

import "testing"

var kfInput = 12
var kfResult = 479001600

func TestFactorial(t *testing.T) {
	result := factorial(kfInput)
	if result != kfResult {
		t.Errorf("Factorial was incorrect, input: %d, want: %d, got: %d.", kfInput, kfResult, result)
	}
}
