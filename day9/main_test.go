package main

import "testing"

func TestGroups(t *testing.T) {

	var scoring = []struct {
		line     string
		scoreExpected int
		garbagedExpected int
	}{
		{ "{}", 1, 0 },
		{"{{{}}}", 6, 0},
		{ "{{},{}}", 5, 0},
		{ "{{{},{},{{}}}}", 16, 0},
		{ "{<a>,<a>,<a>,<a>}", 1, 4},
		{ "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{ "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{ "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
		{ "{{<a!>},{<a>},{<a!>},{<ab>}}", 5, 13},
	}
	for _, s := range scoring {
		scoreActual, garbagedActual := scoresAndGarbage(s.line)
		if scoreActual != s.scoreExpected {
			t.Errorf("(S) Line:[%v] expected %v, actual %v", s.line, s.scoreExpected, scoreActual)
		}
		if garbagedActual != s.garbagedExpected {
			t.Errorf("(G) Line:[%v] expected %v, actual %v", s.line, s.garbagedExpected, garbagedActual)
		}
	}

}
