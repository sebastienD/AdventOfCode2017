package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type layer struct {
	index    int
	length   int
	position int
	up       bool
}

func (l *layer) next() {
	if l.up {
		l.position--
		if l.position < 1 {
			l.position = 2
			l.up = false
		}
	} else {
		l.position++
		if l.position > l.length {
			l.position = l.length - 1
			l.up = true
		}
	}
}

func (l *layer) isUp() bool {
	return l.position == 1
}

func (l *layer) severity() int {
	return l.index * l.length
}

func (l *layer) String() string {
	return fmt.Sprintf("i:%v l:%v p:%v u:%v", l.index, l.length, l.position, l.up)
}

func main() {
	file, err := os.Open("input-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	myPosition := 0
	severity := 0
	maxIndex := 0

	firewall := make(map[int]*layer)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ":")
		index, _ := strconv.Atoi(words[0])
		length, _ := strconv.Atoi(strings.Trim(words[1], " "))
		firewall[index] = &layer{index, length, 1, false}
		maxIndex = index
	}

	for myPosition < maxIndex+1 {
		//fmt.Printf("start\t-> %v, severity: %v, position:%v\n", firewall, severity, myPosition)
		l := firewall[myPosition]
		if l != nil && l.isUp() {
			severity += l.severity()
			//fmt.Printf("sev added: %v\n", severity)
		}
		for _, v := range firewall {
			v.next()
		}
		myPosition++
	}
	//fmt.Printf("%v\n", firewall)

	fmt.Printf("Severity %v\n", severity)
}
