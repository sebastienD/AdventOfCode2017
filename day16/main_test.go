package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	   {""},
	}
	for _,p := range tests {
		t.Run(p.name, func(t *testing.T) {
			main()
		})
	}
}
