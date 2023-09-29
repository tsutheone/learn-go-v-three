package main

import "testing"

var ki1 = 42
var ki2 = 16
var r1 = 58
var r2 = 26

func TestAdd(t *testing.T) {
	r := add(ki1, ki2)
	if r != r1 {
		t.Errorf("Add: inputs: %d, %d want: %d, got: %d.", ki1, ki2, r1, r)

	}
}

func TestSubtract(t *testing.T) {
	r := substract(ki1, ki2)
	if r != r2 {
		t.Errorf("Substract: inputs: %d, %d want: %d, got: %d.", ki1, ki2, r1, r)

	}
}

func TestDoMath(t *testing.T) {
	ra, rs := doMath(ki1, ki2, add), doMath(ki1, ki2, substract)
	if ra != r1 {
		t.Errorf("doMath(add): inputs: %d, %d want: %d, got: %d.", ki1, ki2, r1, ra)
	} else if rs != r2 {
		t.Errorf("doMath(substract): inputs: %d, %d want: %d, got: %d.", ki1, ki2, r1, rs)
	}
}
