package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

type program struct {
	name string
}

func main() {


	//programs := make([]program, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	words := strings.Split(scanner.Text(), ",")
	for _,w := range words {
		fmt.Printf("word %v\n", w)
	}

	s := "a"
	for i,v := range s {
		fmt.Printf("value %v %v %v\n", i,v, string(s))
	}
	o := s[0]

	fmt.Printf("End\n")
}