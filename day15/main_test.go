package main

import "testing"

func TestGenerator(t *testing.T) {
	result := generator(5000000, param{16807, 65}, param{48271, 8921})
	if result != 309 {
		t.Fatalf("Expected %v but got %v", 588, result)
	}
}
