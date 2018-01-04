package main

import (
	"reflect"
	"testing"
)

func Test_countAround(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zero", args{"..."}, 0},
		{"One", args{"#.."}, 1},
		{"OneOther", args{"..#"}, 1},
		{"Two", args{"#.#"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAround(tt.args.v); got != tt.want {
				t.Errorf("countAround() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMiddle(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zero", args{"#.."}, 0},
		{"One", args{".##"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMiddle(tt.args.v); got != tt.want {
				t.Errorf("countMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Hash(t *testing.T) {
	type args struct {
		tab []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"02", args{[]string{"..", "##"}}, "02"},
		{"102121101", args{[]string{"..#", ".##", "###"}}, "102121101"},
		{"12", args{[]string{"#.", "##"}}, "12"},
		{"011020100", args{[]string{".#.", "...", "#.#"}}, "011020100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.tab); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestRotate(t *testing.T) {
	type args struct {
		tab []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"2x2", args{[]string{"..", "##"}}, []string{"#.","#."}},
		{"3x3", args{[]string{"...", "###", "##."}}, []string{"##.","##.", ".#."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rotate(tt.args.tab); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

