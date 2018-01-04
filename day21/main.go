package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"regexp"
)

type referential struct {
	ref2 map[string][]string
	ref3 map[string][]string
}

func (r referential) convert(t []string) []string {
	tab := t
	if len(tab) == 2 {
		h := hash2(tab)
		matrix, ok := r.ref2[h]
		for ;!ok; matrix, ok = r.ref2[h] {
			h = hash2(rotate2(tab))
		}
		return matrix
	}
	h := hash3(tab)
	matrix, ok := r.ref3[h]
	for ;!ok; matrix, ok = r.ref3[h] {
		h = hash3(rotate3(tab))
	}
	return matrix
}

func main() {
	ref, err := referentialFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create referential %v \n", err)
	}

	fmt.Println("referential 2", len(ref.ref2), ref.ref2)
	fmt.Println("referential 3", len(ref.ref3), ref.ref3)
}

func referentialFromFile(path string) (referential, error) {
	file, err := os.Open(path)
	if err != nil {
		return referential{}, fmt.Errorf("Open file failed: %v", err)
	}
	defer file.Close()

	ref2 := make(map[string][]string, 0)
	ref3 := make(map[string][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Replace(scanner.Text(), " => ", "/", -1)
		words := strings.Split(line, "/")
		if len(words) == 5 {
			key := []string{words[0], words[1]}
			matrix := []string{words[2], words[3], words[4]}
			ref2[Hash(key)] = matrix
		} else {
			key := []string{words[0], words[1], words[2]}
			matrix := []string{words[3], words[4], words[5]}
			ref3[Hash(key)] = matrix
		}
	}
	return referential{ref2, ref3}, nil
}

func concat(m [][]string) []string {
	return []string{}
}

func split(t []string) [][]string {

	if len(t)%2 == 0 {
		tab := make([][]string, len(t)/2)
		for line:=0;line<len(t)/2;line++ {
			rex := regexp.MustCompile(".{2}")
			rex.Split(t[line], -1)

			for sub := 0; sub < len(t)/2; sub++ {

			}
		}

		tab[0] =
		tab[1] =
		return tab
	}
	tab := make([][]string, len(t)/3)

	return tab
}