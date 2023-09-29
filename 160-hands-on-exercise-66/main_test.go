package main

import "testing"

func TestRangerAdd(t *testing.T) {
	xi := []string{"Green"}
	if _, got := rangerAdd("Red", xi); got != nil {
		t.Fatalf("rangerAdd: %v\n", got)
	}
}
