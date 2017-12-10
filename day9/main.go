package main

import (
	"fmt"
	"regexp"
	"os"
	"github.com/mgutz/logxi/v1"
	"bufio"
)


func scoresAndGarbage(line string) (int,int) {
	r := regexp.MustCompile("!.{1}")
	line = r.ReplaceAllString(line, "")

	r = regexp.MustCompile("<[^>]*>")
	trash := r.FindAllString(line, -1)
	garbaged := 0
	for _,c := range trash {
		//fmt.Printf("--- %v\n", c)
		garbaged += len(c)-2
	}
	line = r.ReplaceAllString(line, "")

	//fmt.Printf("%v\n", line)
	score := 0
	level := 0
	for _,c := range line {
		if c == '{' {
			level++
			score += level
		}
		if c == '}' {
			level--
		}
	}
	return score, garbaged
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	score, garbaged  := scoresAndGarbage(line)


	fmt.Printf("score:%v garbaged:%v\n", score, garbaged)
}
