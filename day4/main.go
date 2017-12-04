package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countGood := 0
	countGoodAnag := 0
	countAll := 0
	for scanner.Scan() {
		line := scanner.Text()
		words := make(map[string]string)
		scan := bufio.NewScanner(strings.NewReader(line))
		scan.Split(bufio.ScanWords)
		count := 0
		for scan.Scan() {
			count++
			words[scan.Text()] = ""
		}
		if count == len(words) {
			countGood++
			anags := make(map[string]string)
			for k := range words {
				anags[normalize(k)] = ""
			}
			if len(anags) == len(words) {
				countGoodAnag++
			}
		}
		countAll++
	}

	fmt.Printf("Nombre de ligne OK: %v/%d\n", countGood, countAll)
	fmt.Printf("Nombre de ligne OK Anag: %v/%d\n", countGoodAnag, countGood)
}

func normalize(word string) string {
	runes := []string{}
	for _, Rune := range word {
		runes = append(runes, string(Rune))
	}
	sort.Strings(runes)
	return strings.Join(runes, "")
}
