package main

import (
	"testing"
	"strings"
)

func Test_main(t *testing.T) {
	tests := []struct {
		word string
		path string
		steps int
	}{
	{"ABCDEF", `     |
     |  +--+
     A  |  C
 F---|----E|--+
     |  |  |  D
     +B-+  +--+
`, 38},
	}
	for _,tt := range tests {
		m := strings.Split(tt.path, "\n")

		c := newCard(m)
		for c.next() {}

		if c.word() != tt.word {
			t.Fatalf("Word expected is %v but was %v\n", tt.word, c.word())
		}
		if c.steps != tt.steps {
			t.Fatalf("Steps expected is %v but was %v\n", tt.steps, c.steps)
		}
	}
}
