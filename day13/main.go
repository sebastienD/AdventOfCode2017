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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wait := 0

	firewall := make(map[int]*layer)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), ":")
		index, _ := strconv.Atoi(words[0])
		length, _ := strconv.Atoi(strings.Trim(words[1], " "))
		firewall[index] = &layer{index, length, 1, false}
	}

	for !simulation2(wait, firewall) {
		wait++
	}

	fmt.Printf("Wait %v\n", wait)
}

func simulation2(wait int, firewall map[int]*layer) bool {
	//fmt.Printf("simu with wait %v\n", wait)
	for k, v := range firewall {
		if (wait+k)%((v.length-1)*2) == 0 {
			//fmt.Printf("KO wait:%v index:%v l:%v\n", wait, k, v.length)
			return false
		}
	}
	return true
}

func simulation(wait int, firewall map[int]*layer) int {
	//fmt.Printf("simu with wait %v\n", wait)
	for wait > 0 {
		for _, v := range firewall {
			v.next()
		}
		wait--
	}
	myPosition := 0
	for goForward(myPosition, firewall) {
		myPosition++
	}
	return myPosition
}

func goForward(position int, firewall map[int]*layer) bool {
	l := firewall[position]
	if l != nil && l.isUp() {
		//fmt.Printf("KO with position %v\n", position)
		return false
	}
	for _, v := range firewall {
		v.next()
	}
	return true
}

func copyFirewall(firewall map[int]*layer) map[int]*layer {
	f := make(map[int]*layer)
	for k, v := range firewall {
		f[k] = &layer{v.index, v.length, v.position, v.up}
	}
	return f
}
