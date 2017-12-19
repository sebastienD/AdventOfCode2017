package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"unicode"
	"regexp"
)

type direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type point struct {
	x int
	y int
}

type card struct {
	p *point
	m []string
	d direction
	letters []string
	steps int
}

func newCard(m []string) *card {
	x := strings.Index(m[0], "|")
	return &card{ &point{x, 0}, m, DOWN, make([]string, 0), 0}
}

func (c *card) checkLetter() {
	s := int32(c.m[c.p.y][c.p.x])
	if unicode.IsLetter(s) {
		c.letters = append(c.letters, string(s))
	}
}

func (c *card) isValid(p point) bool {
	if p.y < 0 || p.y >= len(c.m) {
		return false
	}
	line := c.m[p.y]
	if p.x < 0 || p.x >= len(line) {
		return false
	}
	re := regexp.MustCompile("[A-Z\\-_|+]")
	return re.Match([]byte{c.m[p.y][p.x]})
}

func (c *card)next() bool {
	fmt.Println("direction", c.d, string(c.m[c.p.y][c.p.x]))
	c.steps++
	switch c.d {
	case DOWN:
		if c.isValid(point{c.p.x, c.p.y+1}) {
			c.p.y++
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x+1, c.p.y}) {
			c.p.x++
			c.d = RIGHT
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x-1, c.p.y}) {
			c.p.x--
			c.d = LEFT
			c.checkLetter()
			return true
		}
	case UP:
		if c.isValid(point{c.p.x, c.p.y-1}) {
			c.p.y--
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x+1, c.p.y}) {
			c.p.x++
			c.d = RIGHT
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x-1, c.p.y}) {
			c.p.x--
			c.d = LEFT
			c.checkLetter()
			return true
		}
	case RIGHT:
		po := point{c.p.x+1, c.p.y}
		if c.isValid(po) {
			c.p.x++
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x, c.p.y+1}) {
			c.p.y++
			c.d = DOWN
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x, c.p.y-1}) {
			c.p.y--
			c.d = UP
			c.checkLetter()
			return true
		}
	case LEFT:
		po := point{c.p.x-1, c.p.y}
		if c.isValid(po) {
			c.p.x--
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x, c.p.y+1}) {
			c.p.y++
			c.d = DOWN
			c.checkLetter()
			return true
		}
		if c.isValid(point{c.p.x, c.p.y-1}) {
			c.p.y--
			c.d = UP
			c.checkLetter()
			return true
		}
	}
	return false
}

func (c * card) word() string {
	return strings.Join(c.letters, "")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open file failed")
	}
	defer file.Close()

	m := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m = append(m, scanner.Text())
	}

	card := newCard(m)
	for card.next() {}

	fmt.Println("words", card.word(), "steps", card.steps)
}
