package main

import "testing"

func Test_registers_set(t *testing.T) {
	steps := []struct {
		name string
		varVal string
		expected int64
	}{
		{"a", "1", 1},
		{"b", "p", 0},
	}
	for _,s := range steps {
		registers := newRegisters()
		registers.set(s.name, s.varVal)
		if registers.reg[s.name] != s.expected {
			t.Fatalf("Expected %v but got %v", s.expected, registers.reg[s.name])
		}
	}
}

func Test_registers_add(t *testing.T) {
	steps := []struct {
		name string
		varVal string
		expected int64
	}{
		{"a", "1", 1},
		{"b", "p", 0},
	}
	for _,s := range steps {
		registers := newRegisters()
		registers.add(s.name, s.varVal)
		if registers.reg[s.name] != s.expected {
			t.Fatalf("Expected %v but got %v", s.expected, registers.reg[s.name])
		}
	}
}

func Test_registers_mul(t *testing.T) {
	steps := []struct {
		name string
		varVal string
		expected int64
	}{
		{"a", "a", 0},
		{"b", "p", 0},
	}
	for _,s := range steps {
		registers := newRegisters()
		registers.mul(s.name, s.varVal)
		if registers.reg[s.name] != s.expected {
			t.Fatalf("Expected %v but got %v", s.expected, registers.reg[s.name])
		}
	}
}

func Test_registers_mod(t *testing.T) {
	steps := []struct {
		name string
		varVal string
		expected int64
	}{
		{"a", "3", 0},
		{"a", "5", 4},
	}
	for _,s := range steps {
		registers := newRegisters()
		registers.set(s.name, "9")
		registers.mod(s.name, s.varVal)
		if registers.reg[s.name] != s.expected {
			t.Fatalf("Expected %v but got %v", s.expected, registers.reg[s.name])
		}
	}
}
