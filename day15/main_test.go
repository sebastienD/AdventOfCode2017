package main

import "testing"

func TestGenerator(t *testing.T) {
	result := generator(40000000, param{16807, 65}, param{48271, 8921})
	if result != 588 {
		t.Fatalf("Expected %s but got %s", 588, result)
	}
}
